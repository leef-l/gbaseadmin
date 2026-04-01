import { useState, useRef, useCallback } from 'react';
import { View, Text, Input } from '@tarojs/components';
import Taro from '@tarojs/taro';
import { sendCode, login, wxLogin } from '../../api/auth';
import { getMemberInfo } from '../../api/member';
import { useAuthStore } from '../../store/auth';
import './index.scss';

export default function LoginPage() {
  const { setToken, setUserInfo } = useAuthStore();
  const [phone, setPhone] = useState('');
  const [code, setCode] = useState('');
  const [agreed, setAgreed] = useState(false);
  const [countdown, setCountdown] = useState(0);
  const [loading, setLoading] = useState(false);
  const timerRef = useRef<any>(null);

  const startCountdown = useCallback(() => {
    setCountdown(60);
    timerRef.current = setInterval(() => {
      setCountdown((prev) => {
        if (prev <= 1) {
          clearInterval(timerRef.current);
          return 0;
        }
        return prev - 1;
      });
    }, 1000);
  }, []);

  const handleSendCode = async () => {
    if (countdown > 0) return;
    if (!phone || phone.length !== 11) {
      Taro.showToast({ title: '请输入正确的手机号', icon: 'none' });
      return;
    }
    try {
      await sendCode(phone);
      Taro.showToast({ title: '验证码已发送', icon: 'success' });
      startCountdown();
    } catch (e) {
      Taro.showToast({ title: '发送失败', icon: 'none' });
    }
  };

  const handleLogin = async () => {
    if (loading) return;
    if (!phone || phone.length !== 11) {
      Taro.showToast({ title: '请输入正确的手机号', icon: 'none' });
      return;
    }
    if (!code) {
      Taro.showToast({ title: '请输入验证码', icon: 'none' });
      return;
    }
    if (!agreed) {
      Taro.showToast({ title: '请同意用户协议', icon: 'none' });
      return;
    }
    setLoading(true);
    try {
      const res = await login(phone, code);
      if (res?.token) {
        setToken(res.token);
        try {
          const info = await getMemberInfo();
          if (info) setUserInfo(info);
        } catch (_) {
          // 获取用户信息失败不阻塞登录流程
        }
        Taro.showToast({ title: '登录成功', icon: 'success' });
        setTimeout(() => Taro.navigateBack(), 1000);
      }
    } catch (e) {
      Taro.showToast({ title: '登录失败', icon: 'none' });
    } finally {
      setLoading(false);
    }
  };

  const handleWxLogin = async () => {
    const env = Taro.getEnv();
    if (env === Taro.ENV_TYPE.WEAPP) {
      // 小程序环境：调用 wx.login 获取 code
      try {
        const loginRes = await Taro.login();
        const res = await wxLogin(loginRes.code);
        if (res?.token) {
          setToken(res.token);
          try {
            const info = await getMemberInfo();
            if (info) setUserInfo(info);
          } catch (_) {}
          Taro.showToast({ title: '登录成功', icon: 'success' });
          setTimeout(() => Taro.navigateBack(), 1000);
        }
      } catch (e) {
        Taro.showToast({ title: '微信登录失败', icon: 'none' });
      }
    } else {
      // H5 环境
      Taro.showToast({ title: '请在微信中打开', icon: 'none' });
    }
  };

  const handleAlipayLogin = () => {
    Taro.showToast({ title: '支付宝登录即将开放', icon: 'none' });
  };

  return (
    <View className="login">
      <View className="login__header">
        <Text className="login__logo">🎮</Text>
        <Text className="login__app-name">陪玩平台</Text>
      </View>

      <View className="login__form">
        <View className="login__field">
          <Input
            className="login__input"
            type="number"
            placeholder="请输入手机号"
            maxlength={11}
            value={phone}
            onInput={(e) => setPhone(e.detail.value)}
          />
        </View>
        <View className="login__field login__field--code">
          <Input
            className="login__input"
            type="number"
            placeholder="请输入验证码"
            maxlength={6}
            value={code}
            onInput={(e) => setCode(e.detail.value)}
          />
          <View
            className={`login__code-btn ${countdown > 0 ? 'login__code-btn--disabled' : ''}`}
            onClick={handleSendCode}
          >
            <Text>{countdown > 0 ? `${countdown}s` : '获取验证码'}</Text>
          </View>
        </View>

        <View
          className={`login__btn ${loading ? 'login__btn--disabled' : ''}`}
          onClick={handleLogin}
        >
          <Text>{loading ? '登录中...' : '登录'}</Text>
        </View>

        <View className="login__agreement" onClick={() => setAgreed(!agreed)}>
          <View className={`login__checkbox ${agreed ? 'login__checkbox--checked' : ''}`}>
            {agreed && <Text>✓</Text>}
          </View>
          <Text className="login__agreement-text">
            我已阅读并同意《用户协议》和《隐私政策》
          </Text>
        </View>

        {/* 第三方登录 */}
        <View className="login__divider">
          <View className="login__divider-line" />
          <Text className="login__divider-text">其他登录方式</Text>
          <View className="login__divider-line" />
        </View>

        <View className="login__third-party">
          <View className="login__third-item" onClick={handleWxLogin}>
            <View className="login__third-icon login__third-icon--wx">
              <Text className="login__third-icon-text">微</Text>
            </View>
            <Text className="login__third-label">微信登录</Text>
          </View>
          <View className="login__third-item" onClick={handleAlipayLogin}>
            <View className="login__third-icon login__third-icon--alipay">
              <Text className="login__third-icon-text">支</Text>
            </View>
            <Text className="login__third-label">支付宝登录</Text>
          </View>
        </View>
      </View>
    </View>
  );
}
