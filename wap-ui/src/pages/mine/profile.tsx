import { useState } from 'react';
import { View, Text, Input, Picker } from '@tarojs/components';
import Taro, { useLoad } from '@tarojs/taro';
import { updateMember } from '../../api/member';
import { useAuthStore } from '../../store/auth';
import './profile.scss';

const genderOptions = ['保密', '男', '女'];

export default function ProfilePage() {
  const { userInfo, setUserInfo } = useAuthStore();
  const [nickname, setNickname] = useState('');
  const [gender, setGender] = useState(0);
  const [birthday, setBirthday] = useState('');
  const [avatar, setAvatar] = useState('');
  const [saving, setSaving] = useState(false);

  useLoad(() => {
    if (userInfo) {
      setNickname(userInfo.nickname || '');
      setAvatar(userInfo.avatar || '');
    }
  });

  const handleChooseAvatar = () => {
    Taro.chooseImage({
      count: 1,
      sizeType: ['compressed'],
      sourceType: ['album', 'camera'],
      success: (res) => {
        setAvatar(res.tempFilePaths[0]);
      },
    });
  };

  const handleSave = async () => {
    if (saving) return;
    if (!nickname.trim()) {
      Taro.showToast({ title: '请输入昵称', icon: 'none' });
      return;
    }
    setSaving(true);
    try {
      await updateMember({ nickname, gender, birthday, avatar });
      setUserInfo({ ...userInfo!, nickname, avatar });
      Taro.showToast({ title: '保存成功', icon: 'success' });
      setTimeout(() => Taro.navigateBack(), 1000);
    } catch (e) {
      Taro.showToast({ title: '保存失败', icon: 'none' });
    } finally {
      setSaving(false);
    }
  };

  return (
    <View className="profile">
      <View className="profile__avatar-wrap" onClick={handleChooseAvatar}>
        <View className="profile__avatar">
          {avatar ? '👤' : '📷'}
        </View>
        <Text className="profile__avatar-tip">点击更换头像</Text>
      </View>

      <View className="profile__form card">
        <View className="profile__item">
          <Text className="profile__label">昵称</Text>
          <Input
            className="profile__input"
            placeholder="请输入昵称"
            value={nickname}
            onInput={(e) => setNickname(e.detail.value)}
          />
        </View>
        <View className="profile__item">
          <Text className="profile__label">性别</Text>
          <Picker mode="selector" range={genderOptions} value={gender} onChange={(e) => setGender(Number(e.detail.value))}>
            <Text className="profile__value">{genderOptions[gender]} &gt;</Text>
          </Picker>
        </View>
        <View className="profile__item">
          <Text className="profile__label">生日</Text>
          <Picker mode="date" value={birthday} onChange={(e) => setBirthday(e.detail.value)}>
            <Text className="profile__value">{birthday || '请选择'} &gt;</Text>
          </Picker>
        </View>
      </View>

      <View className="profile__btn-wrap">
        <View
          className={`profile__btn ${saving ? 'profile__btn--disabled' : ''}`}
          onClick={handleSave}
        >
          <Text>{saving ? '保存中...' : '保存'}</Text>
        </View>
      </View>
    </View>
  );
}
