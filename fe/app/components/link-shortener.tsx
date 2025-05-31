import { Input, Form } from '@brifui/components';
import { useForm } from 'react-hook-form';

type FormValues = {
  url: string;
};

export const LinkShortener = () => {
  const { register, handleSubmit, formState } = useForm<FormValues>();

  const onSubmit = (values: FormValues) => {
    console.log(values);
  };

  return (
    <Form.Root onSubmit={handleSubmit(onSubmit)}>
      <Form.Field name="url" invalid={!!formState.errors.url}>
        <Form.Control>
          <Input
            error={!!formState.errors.url}
            {...register('url', {
              required: { message: 'This field is required', value: true },
            })}
            autoFocus
            placeholder="Paste a link and press Enter to get the short link..."
          />
        </Form.Control>
        <Form.ErrorMessage>{formState.errors.url?.message}</Form.ErrorMessage>
      </Form.Field>
    </Form.Root>
  );
};
