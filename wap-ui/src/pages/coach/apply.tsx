import { useState, useEffect } from 'react';
import { View, Text, Input, Textarea, Picker } from '@tarojs/components';
import Taro from '@tarojs/taro';
import { applyCoach } from '../../api/coach';
import './apply.scss';

export default function CoachApplyPage() {
  const [form, setForm] = useState({
    realName: '',
    phone: '',
    introduction: '',
    skills: '',
    storeId: '',
  });
  const [stores, setStores] = useState<{ id: string; name: string }[]>([]);
  const [storeIndex, setStoreIndex] = useState(-1);
  const [submitting, setSubmitting] = useState(false);

  useEffect(() => {
    // TODO: 接入真实店铺列表API
    setStores([
      { id: '1', name: '旗舰店' },
      { id: '2', name: '大学城店' },
      { id: '3', name: '万达店' },
    ]);
  }, []);

  const handleSubmit = async () => {
    if (!form.realName.trim()) {
      Taro.showToast({ title: '请输入真实姓名', icon: 'none' });
      return;
    }
    if (!form.phone.trim() || !/^1\d{10}$/.test(form.phone)) {
      Taro.showToast({ title: '请输入正确的手机号', icon: 'none' });
      return;
    }
    if (!form.introduction.trim()) {
      Taro.showToast({ title: '请填写自我介绍', icon: 'none' });
      return;
    }
    if (!form.skills.trim()) {
      Taro.showToast({ title: '请填写擅长游戏/技能', icon: 'none' });
      return;
    }
    setSubmitting(true);
    try {
      await applyCoach(form);
      Taro.showToast({ title: '提交成功', icon: 'success' });
      setTimeout(() => Taro.redirectTo({ url: '/pages/coach/apply-status' }), 1500);
    } catch {
      // request 已处理错误提示
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <View className="coach-apply">
      <View className="coach-apply__form">
        <View className="coach-apply__field">
          <Text className="coach-apply__label"><Text className="required">*</Text> 真实姓名</Text>
          <Input className="coach-apply__input" placeholder="请输入真实姓名" value={form.realName} onInput={(e) => setForm({ ...form, realName: e.detail.value })} />
        </View>
        <View className="coach-apply__field">
          <Text className="coach-apply__label"><Text className="required">*</Text> 手机号</Text>
          <Input className="coach-apply__input" type="number" placeholder="请输入手机号" value={form.phone} onInput={(e) => setForm({ ...form, phone: e.detail.value })} maxlength={11} />
        </View>
        <View className="coach-apply__field">
          <Text className="coach-apply__label"><Text className="required">*</Text> 自我介绍</Text>
          <Textarea className="coach-apply__textarea" placeholder="介绍一下自己，让用户更了解你..." value={form.introduction} onInput={(e) => setForm({ ...form, introduction: e.detail.value })} maxlength={500} />
        </View>
        <View className="coach-apply__field">
          <Text className="coach-apply__label"><Text className="required">*</Text> 擅长游戏/技能</Text>
          <Textarea className="coach-apply__textarea" placeholder="如：王者荣耀、英雄联盟、和平精英..." value={form.skills} onInput={(e) => setForm({ ...form, skills: e.detail.value })} maxlength={200} />
        </View>
        <View className="coach-apply__field">
          <Text className="coach-apply__label">所属店铺</Text>
          <Picker mode="selector" range={stores.map(s => s.name)} value={storeIndex} onChange={(e) => {
            const idx = Number(e.detail.value);
            setStoreIndex(idx);
            setForm({ ...form, storeId: stores[idx].id });
          }}>
            <View className="coach-apply__input coach-apply__picker">
              {storeIndex >= 0 ? stores[storeIndex].name : '请选择所属店铺（可选）'}
            </View>
          </Picker>
        </View>
      </View>

      <View className="coach-apply__notice card">
        <Text className="coach-apply__notice-title">申请须知</Text>
        <Text className="coach-apply__notice-text">
          1. 申请人须年满18周岁{'\n'}
          2. 需提供真实有效的个人信息{'\n'}
          3. 审核时间为1-3个工作日{'\n'}
          4. 审核通过后即可开始接单
        </Text>
      </View>

      <View className="bottom-bar">
        <View style={{ flex: 1 }}>
          <View className={`btn-primary ${submitting ? 'btn-disabled' : ''}`} style={{ textAlign: 'center', width: '100%' }} onClick={!submitting ? handleSubmit : undefined}>
            {submitting ? '提交中...' : '提交申请'}
          </View>
        </View>
      </View>
    </View>
  );
}
