import { useState, useRef, useCallback } from 'react';
import { View, Text, Input } from '@tarojs/components';
import Taro from '@tarojs/taro';
import { sendCode, login } from '../../api/auth';
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
      </View>
    </View>
  );
}
