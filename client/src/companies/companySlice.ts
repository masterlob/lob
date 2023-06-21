import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

type Company = {
  id: string;
  name: string;
  description: string;
};

export const commonApi = createApi({
  reducerPath: "commonApi",
  baseQuery: fetchBaseQuery({ baseUrl: "http://localhost:8080/api" }),
  tagTypes: ["Company"],
  endpoints: (builder) => ({
    getCompanyById: builder.query<Company, string>({
      query: (id) => `${id}`,
      providesTags: ["Company"],
    }),
    getCompanies: builder.query<Company[], void>({
      query: () => "companies",
      providesTags: ["Company"],
    }),
    addCompany: builder.mutation<Company, Omit<Company, "id">>({
      query: (body) => ({
        url: "companies",
        method: "POST",
        body,
      }),
      invalidatesTags: ["Company"],
    }),
    editCompany: builder.mutation<
      Company,
      Partial<Company> & Pick<Company, "id">
    >({
      query: (body) => ({
        url: `companies/${body.id}`,
        method: "POST",
        body,
      }),
      invalidatesTags: ["Company"],
    }),
  }),
});

export const {
  useGetCompaniesQuery,
  useGetCompanyByIdQuery,
  useAddCompanyMutation,
  useEditCompanyMutation,
} = commonApi;
