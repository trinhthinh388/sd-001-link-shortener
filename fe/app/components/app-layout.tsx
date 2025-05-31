import type { PropsWithChildren } from 'react';
import { css } from '@brifui/styled/css';
import { Text } from '@brifui/components';

const container = css({
  minW: 'screen',
  minH: 'screen',
  display: 'flex',
  flexDirection: 'column',
  alignItems: 'center',
  justifyContent: 'center',
});

export const AppLayout = ({ children }: PropsWithChildren) => {
  return <div className={container}>{children}</div>;
};
