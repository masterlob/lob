import * as React from "react";
import { useAddCompanyMutation, useGetCompaniesQuery } from "./companySlice";

export default function Company() {
  const { data, error, isLoading } = useGetCompaniesQuery();
  const [addCompany, { isLoading: isUpdating }] = useAddCompanyMutation();

  return (
    <div>
      <div>
        {isUpdating ? (
          <>Adding...</>
        ) : (
          <>
            <button
              onClick={() =>
                addCompany({
                  name: "New Company",
                  description: "My cool new company",
                })
              }
            >
              Add Company
            </button>
          </>
        )}
      </div>

      {error ? (
        <>Oh no, there was an error</>
      ) : isLoading ? (
        <>Loading...</>
      ) : data ? (
        data?.map((company) => (
          <div key={company.id}>
            <h3>{company.name}</h3>
          </div>
        ))
      ) : null}
    </div>
  );
}
