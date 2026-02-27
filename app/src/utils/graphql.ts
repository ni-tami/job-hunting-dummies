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

export const UpdateCompanyMutationTmpl = (
  id: number,
  companyName: string,
  description: string,
  website: string,
) => `mutation updateCompany {
  updateCompany(input: { id: ${id}, website: "${website}", companyName: "${companyName}", description: "${description}" }) {
    companyName
    website
    description
    createdAt
    updatedAt
  }
}
`;

export const UpdateApplicationMutationTmpl = (
  id: number,
  status: string,
) => `mutation updateApplication {
  updateApplication(input: { id: ${id}, status: "${status}" }) {
    applicant {
      user {
        name
        username
      }
    }
    job {
      title
      company {
        companyName
      }
      description
    }
    status
    createdAt
    updatedAt
  }
}
`;

export const UpdateJobMutationTmpl = (
  id: number,
  title: string,
  description: string,
  requirements: string[],
) => `mutation updateJob {
  updateJob(input: { id: ${id}, title: "${title}", description: "${description}", requirements: "${requirements}" }) {
    title
    company {
      companyName
    }
    description
    requirements
    createdAt
    updatedAt
}
`;
