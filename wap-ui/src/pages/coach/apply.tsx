import { useState } from 'react';
import { View, Text, Input, Textarea } from '@tarojs/components';
import Taro from '@tarojs/taro';
import { applyCoach } from '../../api/coach';
import './apply.scss';

export default function CoachApplyPage() {
  const [form, setForm] = useState({
    realName: '',
    idCard: '',
    idCardFrontImage: '',
    idCardBackImage: '',
    skillDesc: '',
  });
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async () => {
    if (!form.realName.trim()) {
      Taro.showToast({ title: '请输入真实姓名', icon: 'none' });
      return;
    }
    if (!form.idCard.trim()) {
      Taro.showToast({ title: '请输入身份证号', icon: 'none' });
      return;
    }
    if (!form.skillDesc.trim()) {
      Taro.showToast({ title: '请填写技能描述', icon: 'none' });
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
          <Text className="coach-apply__label"><Text className="required">*</Text> 身份证号</Text>
          <Input className="coach-apply__input" placeholder="请输入身份证号" value={form.idCard} onInput={(e) => setForm({ ...form, idCard: e.detail.value })} maxlength={18} />
        </View>
        <View className="coach-apply__field">
          <Text className="coach-apply__label">身份证正面照</Text>
          <Input className="coach-apply__input" placeholder="请输入图片URL" value={form.idCardFrontImage} onInput={(e) => setForm({ ...form, idCardFrontImage: e.detail.value })} />
        </View>
        <View className="coach-apply__field">
          <Text className="coach-apply__label">身份证反面照</Text>
          <Input className="coach-apply__input" placeholder="请输入图片URL" value={form.idCardBackImage} onInput={(e) => setForm({ ...form, idCardBackImage: e.detail.value })} />
        </View>
        <View className="coach-apply__field">
          <Text className="coach-apply__label"><Text className="required">*</Text> 技能描述</Text>
          <Textarea className="coach-apply__textarea" placeholder="如：王者荣耀、英雄联盟、和平精英..." value={form.skillDesc} onInput={(e) => setForm({ ...form, skillDesc: e.detail.value })} maxlength={500} />
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
