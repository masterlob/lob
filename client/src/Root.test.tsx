import React from 'react';
import { render } from '@testing-library/react';
import Root from './Root';

test('renders login button', () => {
  const {getByRole} = render(<Root />);
  
  const button = getByRole("button");
  expect(button).toBeInTheDocument();
});
