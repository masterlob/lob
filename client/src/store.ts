import { configureStore } from "@reduxjs/toolkit";
import { commonApi } from "./companies/companySlice";
import loginReducer from "./login/loginSlice";

export const store = configureStore({
  reducer: {
    [commonApi.reducerPath]: commonApi.reducer,
    login: loginReducer,
  },

  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(commonApi.middleware),
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
