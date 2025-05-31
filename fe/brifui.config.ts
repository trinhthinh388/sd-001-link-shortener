import { defineConfig } from '@brifui/theme';

export default defineConfig({
  include: [
    './app/**/*.{ts,tsx,js,jsx}',
    './node_modules/@brifui/**/*.{ts,tsx,js,jsx}',
  ],
  theme: {
    extends: {
      tokens: {
        fonts: {
          heading: {
            value: '"Funnel Display", sans-serif',
          },
        },
      },
    },
  },
});
