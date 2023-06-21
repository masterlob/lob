import { createSlice, PayloadAction } from "@reduxjs/toolkit";

// Define a type for the slice state
export interface UserLoginState {
  isLoggedIn: boolean;
}

export interface UserLoginData {
  user: string;
  password: string;
};

// Define the initial state using that type
const initialState: UserLoginState = {
  isLoggedIn: false,
};

export const userLoginSlice = createSlice({
  name: "userLogin",
  initialState,
  reducers: {
    login: (state, action: PayloadAction<UserLoginData>) => {
      // FIXME!!!!!: validate the user name and password by calling the backend
      // store the JWT in a cookie...
      state.isLoggedIn = true;
    },
    logout: (state) => {
      state.isLoggedIn = false;
    },
  },
});

export const { login, logout } = userLoginSlice.actions;
export default userLoginSlice.reducer;
