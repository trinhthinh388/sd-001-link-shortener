import { Card, Text } from '@brifui/components';
import { css } from '@brifui/styled/css';
import { LinkShortener } from '~/components';

import type { Route } from './+types/home';

export function meta({}: Route.MetaArgs) {
  return [
    { title: 'Pocketlink | The shorten URL application' },
    { name: 'description', content: 'Welcome to Pocketlink!' },
  ];
}

const card = css.raw({
  width: '100%',
  maxWidth: 'container.md',
});

const cardHeader = css.raw({
  textAlign: 'center',
});

const cardBody = css.raw({
  mb: 4,
});

const title = css.raw({
  marginBottom: '6',
});

export default function Home() {
  return (
    <Card css={card}>
      <Card.Header css={cardHeader}>
        <Text css={title} type="heading.5xl">
          Pocketlink
        </Text>
      </Card.Header>
      <Card.Body css={cardBody}>
        <LinkShortener />
      </Card.Body>
    </Card>
  );
}
