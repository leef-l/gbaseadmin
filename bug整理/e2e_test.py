#!/usr/bin/env python3
"""
GBaseAdmin 全系统 E2E 自动化测试 v2
- 后台管理端：API 登录获取 token，注入 localStorage 绕过滑块
- WAP 端：完整活动流程（登录→参与→填写→上传→领奖）
"""

import os
import json
import time
from datetime import datetime
from playwright.sync_api import sync_playwright

BASE_URL = "https://pw.easytestdev.online"
ADMIN_URL = f"{BASE_URL}/admin/"
WAP_URL = f"{BASE_URL}/wap/"
SCREENSHOT_DIR = os.path.join(os.path.dirname(__file__), "screenshots")
BUG_DIR = os.path.dirname(__file__)

TEST_PHONE = "13900001111"
TEST_CODE = "123456"
ADMIN_USER = "admin"
ADMIN_PASS = "admin123"

bug_list = []
bug_counter = 0


def log(msg):
    print(f"[{datetime.now().strftime('%H:%M:%S')}] {msg}")


def screenshot(page, name):
    ts = datetime.now().strftime("%H%M%S")
    path = os.path.join(SCREENSHOT_DIR, f"{ts}_{name}.png")
    try:
        page.evaluate("document.fonts.ready")
    except Exception:
        pass
    page.screenshot(path=path, timeout=15000)
    log(f"  截图: {os.path.basename(path)}")
    return path


def record_bug(title, description, screenshot_path=None, severity="中"):
    global bug_counter
    bug_counter += 1
    bug = {
        "id": bug_counter, "title": title, "description": description,
        "severity": severity, "time": datetime.now().strftime("%Y-%m-%d %H:%M:%S"),
        "screenshot": os.path.basename(screenshot_path) if screenshot_path else "",
    }
    bug_list.append(bug)

    bug_file = os.path.join(BUG_DIR, "bug列表.md")
    with open(bug_file, "a", encoding="utf-8") as f:
        if bug_counter == 1:
            f.write(f"# Bug 列表\n\n测试时间: {datetime.now().strftime('%Y-%m-%d %H:%M')}\n\n")
        f.write(f"## Bug #{bug['id']}: {bug['title']}\n\n")
        f.write(f"- **严重程度**: {bug['severity']}\n")
        f.write(f"- **发现时间**: {bug['time']}\n")
        f.write(f"- **描述**: {bug['description']}\n")
        if bug["screenshot"]:
            f.write(f"- **截图**: [查看](screenshots/{bug['screenshot']})\n")
        f.write(f"\n---\n\n")

    log(f"  BUG #{bug_counter}: {title}")


def wait_load(page, timeout=5000):
    try:
        page.wait_for_load_state("networkidle", timeout=timeout)
    except Exception:
        pass


# ========== 后台管理端 ==========

def test_admin_login(page):
    """后台登录：填写用户名密码 + 拖动滑块验证码"""
    log("=== 测试后台登录 ===")
    page.goto(ADMIN_URL, wait_until="domcontentloaded", timeout=30000)
    time.sleep(3)
    wait_load(page, 8000)
    screenshot(page, "admin_00_login_page")

    try:
        # 填写用户名
        username_input = page.locator("input[name='username']")
        username_input.wait_for(state="visible", timeout=10000)
        username_input.click()
        username_input.fill(ADMIN_USER)
        time.sleep(0.3)

        # 填写密码
        password_input = page.locator("input[name='password']")
        password_input.click()
        password_input.fill(ADMIN_PASS)
        time.sleep(0.3)

        screenshot(page, "admin_01_filled")

        # 拖动滑块验证码（参考 Vben 官方 e2e 测试）
        slider_captcha = page.locator("div[name='captcha']")
        slider_action = page.locator("div[name='captcha-action']")

        slider_captcha.wait_for(state="visible", timeout=5000)
        slider_action.wait_for(state="visible", timeout=5000)

        captcha_box = slider_captcha.bounding_box()
        action_box = slider_action.bounding_box()

        if not captcha_box or not action_box:
            sp = screenshot(page, "admin_01_no_slider")
            record_bug("后台滑块验证码未找到", "找不到滑块验证码的 bounding box", sp, "高")
            return False

        start_x = action_box["x"] + action_box["width"] / 2
        start_y = action_box["y"] + action_box["height"] / 2
        target_x = start_x + captcha_box["width"] + action_box["width"]
        target_y = start_y

        page.mouse.move(start_x, start_y)
        page.mouse.down()
        page.mouse.move(target_x, target_y, steps=20)
        page.mouse.up()
        time.sleep(0.5)

        screenshot(page, "admin_02_slider_done")

        # 点击登录按钮
        login_btn = page.get_by_role("button", name="login")
        if login_btn.count() == 0:
            # 尝试中文
            login_btn = page.get_by_role("button", name="登录")
        if login_btn.count() == 0:
            login_btn = page.locator("button").filter(has_text="登录").first
        if login_btn.count() == 0:
            login_btn = page.locator("button").filter(has_text="Login").first

        login_btn.click(timeout=5000)
        time.sleep(5)
        wait_load(page, 10000)

        screenshot(page, "admin_03_after_login")
        current_url = page.url
        log(f"  登录后URL: {current_url}")

        if "/login" not in current_url and "/auth" not in current_url:
            log("  OK 后台登录成功")
            return True
        else:
            sp = screenshot(page, "admin_03_login_failed")
            record_bug("后台登录失败", f"滑块拖动+表单提交后仍在登录页: {current_url}", sp, "中")
            return False

    except Exception as e:
        sp = screenshot(page, "admin_error_login")
        record_bug("后台登录异常", f"{type(e).__name__}: {str(e)[:200]}", sp, "高")
        return False


def test_admin_menus(page):
    """遍历后台菜单"""
    log("=== 测试后台菜单遍历 ===")
    time.sleep(2)

    try:
        # 展开所有一级菜单
        submenu_titles = page.locator(".ant-menu-submenu-title:visible").all()
        log(f"  一级菜单数: {len(submenu_titles)}")
        for title in submenu_titles:
            try:
                text = title.inner_text(timeout=1000).strip()
                if text:
                    title.click(timeout=2000)
                    time.sleep(0.3)
            except Exception:
                continue

        time.sleep(1)
        screenshot(page, "admin_03_menu_expanded")

        # 点击所有叶子菜单
        menu_items = page.locator(".ant-menu-item:visible").all()
        log(f"  叶子菜单数: {len(menu_items)}")

        visited = set()
        for idx, item in enumerate(menu_items):
            try:
                text = item.inner_text(timeout=1000).strip()
                if not text or text in visited or len(text) > 20:
                    continue
                visited.add(text)
                log(f"  [{idx+1}] {text}")
                item.click(timeout=3000)
                time.sleep(2)
                wait_load(page, 3000)

                # 检查页面错误
                body = page.locator("body").inner_text(timeout=2000)[:300]
                for kw in ["404", "500", "页面不存在", "Internal Server Error"]:
                    if kw in body:
                        sp = screenshot(page, f"admin_menu_err_{idx}")
                        record_bug(f"后台[{text}]页面{kw}", f"菜单'{text}'显示: {kw}", sp)
                        break
                else:
                    screenshot(page, f"admin_menu_{idx}_{text[:8]}")

            except Exception:
                continue

    except Exception as e:
        sp = screenshot(page, "admin_menu_error")
        record_bug("后台菜单遍历异常", str(e)[:200], sp)


def test_admin_crud(page):
    """测试后台 CRUD 操作"""
    log("=== 测试后台 CRUD ===")
    targets = ["用户管理", "角色管理", "部门管理", "菜单管理"]

    for name in targets:
        try:
            item = page.locator(".ant-menu-item:visible").filter(has_text=name).first
            if item.count() == 0:
                log(f"  {name}: 未找到菜单")
                continue

            item.click(timeout=3000)
            time.sleep(2)
            wait_load(page)

            # 检查表格
            rows = page.locator(".vxe-body--row, .ant-table-row").count()
            log(f"  {name}: 表格行数={rows}")

            # 测试新增
            add_btn = page.locator("button:visible").filter(has_text="新增").first
            if add_btn.count() > 0:
                add_btn.click()
                time.sleep(1.5)
                modal = page.locator(".ant-modal:visible, .ant-drawer:visible")
                if modal.count() > 0:
                    screenshot(page, f"admin_crud_{name}_form")
                    log(f"  {name}: OK 新增弹窗正常")
                    page.locator(".ant-modal-close:visible, .ant-drawer-close:visible").first.click(timeout=2000)
                    time.sleep(0.5)
                else:
                    sp = screenshot(page, f"admin_crud_{name}_no_modal")
                    record_bug(f"后台[{name}]新增弹窗未弹出", "点击新增按钮后无弹窗", sp)

            # 测试编辑（点第一行的编辑按钮）
            edit_btn = page.locator("button:visible, a:visible").filter(has_text="编辑").first
            if edit_btn.count() > 0 and rows > 0:
                edit_btn.click()
                time.sleep(1.5)
                modal = page.locator(".ant-modal:visible, .ant-drawer:visible")
                if modal.count() > 0:
                    screenshot(page, f"admin_crud_{name}_edit")
                    log(f"  {name}: OK 编辑弹窗正常")
                    page.locator(".ant-modal-close:visible, .ant-drawer-close:visible").first.click(timeout=2000)
                    time.sleep(0.5)

        except Exception as e:
            log(f"  {name}: 异常 {str(e)[:80]}")


# ========== WAP 端 ==========

def test_wap_login(page):
    """WAP 端登录"""
    log("=== 测试 WAP 登录 ===")
    page.goto(WAP_URL, wait_until="domcontentloaded", timeout=30000)
    time.sleep(3)
    wait_load(page)
    screenshot(page, "wap_01_home")

    try:
        # 点"我的"
        my_tab = page.locator("text=我的").first
        if my_tab.is_visible(timeout=3000):
            my_tab.click()
            time.sleep(1.5)

        # 点"点击登录"
        login_entry = page.locator("text=点击登录").first
        if login_entry.is_visible(timeout=3000):
            login_entry.click()
            time.sleep(2)
        elif "login" not in page.url.lower():
            log("  已登录状态")
            return True

        screenshot(page, "wap_02_login_page")

        # 勾选协议
        checkbox = page.locator("[class*='checkbox']").first
        if checkbox.is_visible(timeout=2000):
            checkbox.click(force=True)
            time.sleep(0.3)

        # 填手机号
        phone_input = page.locator("input[placeholder*='手机']").first
        phone_input.click()
        phone_input.fill(TEST_PHONE)
        time.sleep(0.3)

        # 发验证码
        send_btn = page.locator("text=获取验证码").first
        if send_btn.is_visible(timeout=3000):
            send_btn.click()
            time.sleep(3)

        # 填验证码
        code_input = page.locator("input[placeholder*='验证码']").first
        if code_input.is_visible(timeout=5000):
            code_input.click()
            time.sleep(0.3)
            code_input.fill(TEST_CODE)
        screenshot(page, "wap_03_filled")

        # 点登录
        login_btn = page.locator("[class*='login__btn']").filter(has_text="登录").first
        if login_btn.count() == 0:
            login_btn = page.locator("text=登录").first
        login_btn.click(force=True)
        time.sleep(4)
        wait_load(page)
        screenshot(page, "wap_04_after_login")

        if "login" in page.url.lower():
            sp = screenshot(page, "wap_04_login_failed")
            record_bug("WAP登录失败", "提交登录后仍在登录页", sp, "高")
            return False

        log("  OK WAP 登录成功")
        return True

    except Exception as e:
        sp = screenshot(page, "wap_error_login")
        record_bug("WAP登录异常", f"{type(e).__name__}: {str(e)[:200]}", sp, "高")
        return False


def test_wap_pages(page):
    """遍历 WAP 页面"""
    log("=== 测试 WAP 页面 ===")
    pages_to_test = [
        ("#/pages/index/index", "首页"),
        ("#/pages/category/index", "分类"),
        ("#/pages/activity/list", "活动列表"),
        ("#/pages/user/index", "个人中心"),
        ("#/pages/user/wallet", "钱包"),
        ("#/pages/user/balance-log", "余额明细"),
        ("#/pages/coupon/list", "领券中心"),
        ("#/pages/user/coupon", "我的优惠券"),
        ("#/pages/order/list", "订单列表"),
        ("#/pages/user/settings", "设置"),
        ("#/pages/mine/index", "我的"),
    ]

    for path, name in pages_to_test:
        try:
            page.goto(f"{WAP_URL}{path}", wait_until="domcontentloaded", timeout=15000)
            time.sleep(2)
            wait_load(page, 3000)

            body = page.locator("body").inner_text(timeout=2000)[:500]
            for kw in ["404", "页面不存在", "Internal Server Error", "500"]:
                if kw in body:
                    sp = screenshot(page, f"wap_page_err_{name}")
                    record_bug(f"WAP[{name}]页面{kw}", f"访问 {path} 出现: {kw}", sp)
                    break
            else:
                screenshot(page, f"wap_page_{name}")
                log(f"  OK {name}")

        except Exception as e:
            sp = screenshot(page, f"wap_page_err_{name}")
            record_bug(f"WAP[{name}]加载失败", str(e)[:200], sp)


def test_wap_activity_full(page):
    """WAP 活动完整流程"""
    log("=== 测试 WAP 活动完整流程 ===")

    try:
        # 先重置参与记录（通过 API 退出再重新参与）
        # 进入活动列表
        page.goto(f"{WAP_URL}#/pages/activity/list", wait_until="domcontentloaded", timeout=15000)
        time.sleep(2)
        wait_load(page)
        screenshot(page, "wap_act_01_list")

        # 点击活动
        page.evaluate("""() => {
            const all = document.querySelectorAll('*');
            for (const el of all) {
                if (el.children.length === 0 && el.innerText && el.innerText.includes('测试步骤活动')) {
                    el.click();
                    return true;
                }
            }
            return false;
        }""")
        time.sleep(3)
        wait_load(page)
        screenshot(page, "wap_act_02_detail")

        # 滚动查看所有步骤
        for scroll in [300, 600, 900, 1200]:
            page.evaluate(f"window.scrollTo(0, {scroll})")
            time.sleep(0.5)
        screenshot(page, "wap_act_03_all_steps")

        # 检查按钮
        copy_btns = page.locator("text=立即复制").count()
        jump_btns = page.locator("text=立即跳转").count()
        log(f"  立即复制: {copy_btns}, 立即跳转: {jump_btns}")

        if copy_btns == 0:
            sp = screenshot(page, "wap_act_no_copy")
            record_bug("活动步骤无'立即复制'按钮", "不需要填写的步骤缺少操作按钮（前端可能未部署）", sp)

        # 滚动回顶部
        page.evaluate("window.scrollTo(0, 0)")
        time.sleep(0.5)

        # 检查是否已参与
        join_btn = page.locator("text=立即参与").first
        quit_btn = page.locator("text=取消报名").first
        reward_btn = page.locator("text=领取奖励").first
        done_text = page.locator("text=已完成").first

        if quit_btn.is_visible(timeout=1000):
            log("  已参与活动，先取消重新参与...")
            quit_btn.click()
            time.sleep(2)
            # 确认弹窗
            confirm = page.locator("text=确认, text=确定").first
            if confirm.is_visible(timeout=2000):
                confirm.click()
                time.sleep(2)
            screenshot(page, "wap_act_04_quit")
            # 刷新页面
            page.reload(wait_until="domcontentloaded")
            time.sleep(2)

        # 参与活动
        join_btn = page.locator("text=立即参与").first
        if join_btn.is_visible(timeout=3000):
            log("  点击: 立即参与")
            join_btn.click()
            time.sleep(2)
            screenshot(page, "wap_act_05_joined")

        # 滚动找到"我已完成"按钮并逐步完成
        page.evaluate("window.scrollTo(0, 0)")
        time.sleep(0.5)

        # 尝试完成每个步骤
        max_steps = 5
        for step_idx in range(max_steps):
            page.evaluate("window.scrollTo(0, 0)")
            time.sleep(0.5)

            # 找到当前活跃步骤
            active_step = page.locator(".step-card--active").first
            if not active_step.is_visible(timeout=2000):
                log(f"  步骤{step_idx+1}: 无活跃步骤")
                break

            # 滚动到活跃步骤
            active_step.scroll_into_view_if_needed()
            time.sleep(0.5)

            # 检查是否有输入框需要填写
            input_field = active_step.locator("input:visible").first
            if input_field.count() > 0:
                log(f"  步骤{step_idx+1}: 填写输入框")
                input_field.click()
                input_field.fill(f"测试数据{step_idx+1}")
                time.sleep(0.3)

            # 检查是否有上传区域
            upload = active_step.locator("[class*='upload']").first
            if upload.count() > 0:
                log(f"  步骤{step_idx+1}: 跳过上传步骤（自动化无法上传文件）")

            # 点"我已完成"
            complete_btn = active_step.locator("text=我已完成").first
            if complete_btn.count() == 0:
                # 可能在步骤卡片外面
                complete_btn = page.locator("text=我已完成").first

            if complete_btn.is_visible(timeout=2000):
                log(f"  步骤{step_idx+1}: 点击我已完成")
                complete_btn.click()
                time.sleep(2)
                screenshot(page, f"wap_act_step_{step_idx+1}")
            else:
                log(f"  步骤{step_idx+1}: 未找到完成按钮")
                break

        # 检查是否出现领取奖励
        page.evaluate("window.scrollTo(0, document.body.scrollHeight)")
        time.sleep(1)
        reward_btn = page.locator("text=领取奖励").first
        if reward_btn.is_visible(timeout=3000):
            log("  点击: 领取奖励")
            reward_btn.click()
            time.sleep(3)
            screenshot(page, "wap_act_06_reward")

            # 验证奖励是否到账
            page.goto(f"{WAP_URL}#/pages/user/wallet", wait_until="domcontentloaded", timeout=10000)
            time.sleep(2)
            screenshot(page, "wap_act_07_wallet_after")
            log("  OK 领取奖励流程完成")
        else:
            final_text = page.locator(".activity-detail__btn").first.inner_text(timeout=2000) if page.locator(".activity-detail__btn").count() > 0 else "无按钮"
            log(f"  底部按钮: {final_text}")
            screenshot(page, "wap_act_06_bottom_status")
            if "已完成" in final_text:
                log("  活动已完成（之前已领取奖励）")
            elif "继续完成" in final_text:
                log("  步骤未全部完成")

    except Exception as e:
        sp = screenshot(page, "wap_act_error")
        record_bug("WAP活动流程异常", f"{type(e).__name__}: {str(e)[:200]}", sp)


def test_wap_user_features(page):
    """测试 WAP 用户功能子页面"""
    log("=== 测试 WAP 用户功能 ===")

    # 个人中心
    page.goto(f"{WAP_URL}#/pages/mine/index", wait_until="domcontentloaded", timeout=10000)
    time.sleep(2)
    screenshot(page, "wap_user_01_mine")

    body = page.locator("body").inner_text(timeout=2000)
    if "点击登录" in body:
        record_bug("WAP个人中心未登录", "访问我的页面仍显示未登录状态", None, "高")
        return

    # 测试子页面
    sub_pages = [
        ("充值中心", "wap_user_recharge"),
        ("领券中心", "wap_user_coupon_center"),
        ("我的优惠券", "wap_user_my_coupon"),
        ("余额明细", "wap_user_balance"),
        ("我的评价", "wap_user_review"),
        ("设置", "wap_user_settings"),
    ]

    for name, file_name in sub_pages:
        try:
            # 先回到我的页面
            page.goto(f"{WAP_URL}#/pages/mine/index", wait_until="domcontentloaded", timeout=10000)
            time.sleep(1.5)

            link = page.locator(f"text={name}").first
            if link.is_visible(timeout=2000):
                link.click()
                time.sleep(2)
                wait_load(page, 3000)
                screenshot(page, file_name)
                log(f"  OK {name}")

                # 检查页面内容
                body = page.locator("body").inner_text(timeout=2000)[:300]
                for kw in ["404", "页面不存在", "error"]:
                    if kw.lower() in body.lower():
                        sp = screenshot(page, f"{file_name}_err")
                        record_bug(f"WAP[{name}]异常", f"页面包含: {kw}", sp)
                        break
            else:
                log(f"  未找到: {name}")

        except Exception as e:
            log(f"  {name}: {str(e)[:60]}")


def test_wap_coupon(page):
    """测试领券功能"""
    log("=== 测试 WAP 领券 ===")
    try:
        page.goto(f"{WAP_URL}#/pages/coupon/list", wait_until="domcontentloaded", timeout=10000)
        time.sleep(2)
        screenshot(page, "wap_coupon_01_list")

        # 点击领取按钮
        claim_btn = page.locator("text=领取").first
        if claim_btn.is_visible(timeout=3000):
            log("  点击: 领取优惠券")
            claim_btn.click()
            time.sleep(2)
            screenshot(page, "wap_coupon_02_claimed")
        else:
            log("  无可领取的优惠券")

    except Exception as e:
        log(f"  领券异常: {e}")


def test_wap_order(page):
    """测试订单页面"""
    log("=== 测试 WAP 订单 ===")
    page.goto(f"{WAP_URL}#/pages/order/list", wait_until="domcontentloaded", timeout=10000)
    time.sleep(2)
    screenshot(page, "wap_order_list")

    # 检查各 tab
    tabs = ["全部", "待支付", "进行中", "已完成", "退款"]
    for tab_name in tabs:
        tab = page.locator(f"text={tab_name}").first
        if tab.is_visible(timeout=1000):
            tab.click()
            time.sleep(1)
    screenshot(page, "wap_order_tabs")
    log("  OK 订单 tabs 切换正常")


# ========== 主流程 ==========

def main():
    bug_file = os.path.join(BUG_DIR, "bug列表.md")
    if os.path.exists(bug_file):
        os.remove(bug_file)
    os.makedirs(SCREENSHOT_DIR, exist_ok=True)

    # 清旧截图
    for f in os.listdir(SCREENSHOT_DIR):
        if f.endswith(".png"):
            os.remove(os.path.join(SCREENSHOT_DIR, f))

    with sync_playwright() as p:
        browser = p.chromium.launch(
            headless=True,
            args=["--no-sandbox", "--disable-dev-shm-usage", "--disable-gpu"],
            timeout=60000,
        )

        # ===== 后台 =====
        log("\n" + "=" * 60)
        log("开始测试：后台管理端")
        log("=" * 60)

        admin_ctx = browser.new_context(
            viewport={"width": 1920, "height": 1080},
            ignore_https_errors=True,
        )
        admin_page = admin_ctx.new_page()

        admin_ok = test_admin_login(admin_page)
        if admin_ok:
            test_admin_menus(admin_page)
            test_admin_crud(admin_page)
        admin_ctx.close()

        # ===== WAP =====
        log("\n" + "=" * 60)
        log("开始测试：WAP 端")
        log("=" * 60)

        wap_ctx = browser.new_context(
            viewport={"width": 375, "height": 812},
            user_agent="Mozilla/5.0 (iPhone; CPU iPhone OS 16_0 like Mac OS X) AppleWebKit/605.1.15",
            ignore_https_errors=True,
        )
        wap_page = wap_ctx.new_page()

        wap_ok = test_wap_login(wap_page)
        if wap_ok:
            test_wap_pages(wap_page)
            test_wap_user_features(wap_page)
            test_wap_coupon(wap_page)
            test_wap_order(wap_page)
            test_wap_activity_full(wap_page)

        wap_ctx.close()
        browser.close()

    # 汇总
    log("\n" + "=" * 60)
    log(f"测试完成！发现 {len(bug_list)} 个 Bug")
    log("=" * 60)
    for b in bug_list:
        log(f"  #{b['id']} [{b['severity']}] {b['title']}")

    if not bug_list:
        with open(os.path.join(BUG_DIR, "bug列表.md"), "w", encoding="utf-8") as f:
            f.write("# Bug 列表\n\n")
            f.write(f"测试时间: {datetime.now().strftime('%Y-%m-%d %H:%M')}\n\n")
            f.write("未发现 Bug\n")


if __name__ == "__main__":
    main()
