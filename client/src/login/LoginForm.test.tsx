import React from 'react';
import { render } from '@testing-library/react';
import LoginForm from './LoginForm';

test('renders login button', () => {
  const {getByRole} = render(<LoginForm />);
  
  const button = getByRole("button");
  expect(button).toBeInTheDocument();
});
