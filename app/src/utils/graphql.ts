export const CreateUserMutationTmpl = (name: string, username: string) => `
mutation createUser {
  createUser(input: { name: "${name}", username: "${username}" }) {
    id
    username
    name
    created_at
  }
}
`;

export const CreateApplicantMutationTmpl = (
  userId: number,
) => `mutation createApplicant {
  createApplicant(input: { userId: ${userId} }) {
    id
    user {
      id
    }
    createdAt
  }
}
`;

export const CreateCompanyMutationTmpl = (
  companyName: string,
  userId: number,
  description: string,
  website: string,
) => `mutation createCompany {
  createCompany(input: { companyName: "${companyName}", userId: ${userId}, website: "${website}", description: "${description}" }) {
    id
    user {
      id
    }
    companyName
    createdAt
  }
}
`;
