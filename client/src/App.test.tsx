import React from 'react';
import { render } from '@testing-library/react';
import App from './App';

test('renders login button', () => {
  const {getByRole} = render(<App />);
  
  const button = getByRole("button");
  expect(button).toBeInTheDocument();
});
