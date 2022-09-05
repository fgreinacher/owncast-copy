import React from 'react';
import { Typography } from 'antd';
import { EditSocialLinks } from '../../components/config/EditSocialLinks';

const { Title } = Typography;

export default function ConfigSocialThings() {
  return (
    <div className="config-social-items">
      <Title>Social Items</Title>

      <EditSocialLinks />
    </div>
  );
}
