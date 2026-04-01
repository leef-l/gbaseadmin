import { useState } from 'react';
import { View, Text, Image, Input, Textarea } from '@tarojs/components';
import Taro from '@tarojs/taro';
import { applyCoach } from '../../api/coach';
import { useAuthStore } from '../../store/auth';
import './apply.scss';

const BASE_URL = process.env.TARO_APP_API || '';

async function uploadImage(filePath: string): Promise<string> {
  const token = useAuthStore.getState().token;
  return new Promise((resolve, reject) => {
    Taro.uploadFile({
      url: `${BASE_URL}/api/upload/uploader/upload`,
      filePath,
      name: 'file',
      header: token ? { Authorization: `Bearer ${token}` } : {},
      success: (res) => {
        try {
          const body = JSON.parse(res.data);
          if (body.code === 0 && body.data?.url) {
            resolve(body.data.url as string);
          } else {
            reject(new Error(body.message || '上传失败'));
          }
        } catch {
          reject(new Error('解析响应失败'));
        }
      },
      fail: (err) => reject(err),
    });
  });
}

async function chooseAndUpload(
  setter: (url: string) => void,
  tempSetter: (path: string) => void,
) {
  try {
    const res = await Taro.chooseImage({
      count: 1,
      sizeType: ['compressed'],
      sourceType: ['album', 'camera'],
    });
    const tempPath = res.tempFilePaths[0];
    tempSetter(tempPath);
    Taro.showLoading({ title: '上传中...' });
    const url = await uploadImage(tempPath);
    setter(url);
    Taro.showToast({ title: '上传成功', icon: 'success' });
  } catch {
    Taro.showToast({ title: '图片上传失败，请重试', icon: 'none' });
  } finally {
    Taro.hideLoading();
  }
}

export default function CoachApplyPage() {
  const [form, setForm] = useState({
    realName: '',
    idCard: '',
    idCardFrontImage: '',
    idCardBackImage: '',
    skillDesc: '',
  });
  const [frontPreview, setFrontPreview] = useState('');
  const [backPreview, setBackPreview] = useState('');
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
          <View
            className="coach-apply__upload"
            onClick={() => chooseAndUpload(
              (url) => setForm((prev) => ({ ...prev, idCardFrontImage: url })),
              setFrontPreview,
            )}
          >
            {frontPreview || form.idCardFrontImage ? (
              <Image
                src={frontPreview || form.idCardFrontImage}
                className="coach-apply__upload-preview"
                mode="aspectFill"
              />
            ) : (
              <View className="coach-apply__upload-placeholder">
                <Text className="coach-apply__upload-icon">+</Text>
                <Text className="coach-apply__upload-tip">点击上传正面照</Text>
              </View>
            )}
          </View>
        </View>
        <View className="coach-apply__field">
          <Text className="coach-apply__label">身份证反面照</Text>
          <View
            className="coach-apply__upload"
            onClick={() => chooseAndUpload(
              (url) => setForm((prev) => ({ ...prev, idCardBackImage: url })),
              setBackPreview,
            )}
          >
            {backPreview || form.idCardBackImage ? (
              <Image
                src={backPreview || form.idCardBackImage}
                className="coach-apply__upload-preview"
                mode="aspectFill"
              />
            ) : (
              <View className="coach-apply__upload-placeholder">
                <Text className="coach-apply__upload-icon">+</Text>
                <Text className="coach-apply__upload-tip">点击上传反面照</Text>
              </View>
            )}
          </View>
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
