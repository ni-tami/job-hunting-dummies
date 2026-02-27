import { useState } from "react";
import {
  Button,
  Stack,
  Field,
  Input,
  Tabs,
  Icon,
  TagsInput,
  Fieldset,
  Container,
  NativeSelect,
  For,
} from "@chakra-ui/react";
import { useForm } from "react-hook-form";
import { HiIdentification, HiOfficeBuilding, HiBriefcase } from "react-icons/hi";
import { UpdateCompanyMutationTmpl, UpdateApplicationMutationTmpl, UpdateJobMutationTmpl } from "@/utils/graphql";


type applicationData = {
  id: number,
  status: string,
};

type companyData = {
  id: number,
  companyName: string,
  website: string,
  description: string,
};

type jobData = {
  id: number,
  title: string,
  description: string,
  requirements: string[],
}

interface FormValues {
  updateData: {
    applicationData: applicationData | undefined,
    companyData: companyData | undefined,
    jobData: jobData | undefined,
  },
};
const SPLIT_REGEX = /[;]/;

const userTabTypes = [
  { label: "Application", value: "application" },
  { label: "Company", value: "company" },
  { label: "Job", value: "job" },
];

const applicationStatusTypes = [
  { label: "Applied", value: "APPLIED" },
  { label: "Reviewed", value: "REVIEWED" },
  { label: "Pending Interview", value: "PENDING_INTERVIEW" },
  { label: "Interview Scheduled", value: "INTERVIEW_SCHEDULED" },
  { label: "Pending Result", value: "PENDING_RESULT" },
  { label: "Accepted", value: "ACCEPTED" },
  { label: "Rejected", value: "REJECTED" },
]

const fetchUpdateCompany = async (userFormData: FormValues) => {
  const companyData = userFormData.updateData.companyData;
  if (!companyData) {
    // FIXME
    console.log("Empty company update data: ", userFormData.updateData);
    return undefined;
  }
  const updateCompanyQuery = UpdateCompanyMutationTmpl(
    companyData.id!,
    companyData.companyName!,
    companyData.description!,
    companyData.website!
  );

  console.log("Updating company:", companyData);
  return fetch("http://localhost:8080/query", {
    method: "POST",
    headers: {
      "Content-type": "application/json",
    },
    body: JSON.stringify({
      query: updateCompanyQuery,
      operationName: "updateCompany"
    }),
  })
    .then((res) => res.json())
    .then((json) => json.data);
}

const fetchUpdateApplication = async (userFormData: FormValues) => {
  const applicationData = userFormData.updateData.applicationData;
  if (!applicationData) {
    // FIXME
    console.log("Empty application update data: ", userFormData.updateData);
    return undefined;
  }
  console.log("Updating application:", applicationData);

  const updateApplicationQuery = UpdateApplicationMutationTmpl(
    applicationData.id!,
    applicationData.status!,
  );
  return fetch("http://localhost:8080/query", {
    method: "POST",
    headers: {
      "Content-type": "application/json",
    },
    body: JSON.stringify({
      query: updateApplicationQuery,
      operationName: "updateApplication",
    }),
  })
  .then((res) => res.json())
  .then((json) => json.data);
}

const fetchUpdateJob = async (userFormData: FormValues) => {
  const jobData = userFormData.updateData.jobData;
  if (!jobData) {
    // FIXME
    console.log("Empty application update data: ", userFormData.updateData);
    return undefined;
  }
  console.log("Updating application:", jobData);

  const updateJobQuery = UpdateJobMutationTmpl(
    jobData.id!,
    jobData.title!,
    jobData.description!,
    jobData.requirements!,
  );
  return fetch("http://localhost:8080/query", {
    method: "POST",
    headers: {
      "Content-type": "application/json",
    },
    body: JSON.stringify({
      query: updateJobQuery,
      operationName: "updateJob",
    }),
  })
    .then((res) => res.json())
    .then((json) => json.data);
};

const UpdatePage = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm<FormValues>({
    defaultValues: {
      updateData: {
        applicationData: {
          id: 0,
          status: "",
        },
        companyData: {
          id: 0,
          companyName: "",
          website: "",
          description: "",
        },
        jobData: {
          id: 0,
          title: "",
          description: "",
          requirements: [],
        },
      },
    },
  });

  const onApplicationSubmit = handleSubmit((formSubmitData: FormValues) => {
    console.log("submitted data: ", formSubmitData);
    fetchUpdateApplication(formSubmitData)
      .then((applicationData) => console.log("applicationData: ", applicationData));
  });

  const onCompanySubmit = handleSubmit((formSubmitData: FormValues) => {
    console.log("submitted data: ", formSubmitData);
    fetchUpdateCompany(formSubmitData)
      .then((companyData) => console.log("companyData: ", companyData),
    );
  });

  const onJobSubmit = handleSubmit((formSubmitData: FormValues) => {
    console.log("submitted data: ", formSubmitData);
    fetchUpdateJob(formSubmitData)
    .then((jobData) => console.log("jobData: ", jobData),
    );
  });

  const [userTabType, setUserTabType] = useState<string>("application");
  const onTabSwitched = (details: Tabs.TabsValueChangeDetails) => {
    // TODO: may be related to controller, do not reset if partially filled
    reset();
    setUserTabType(details.value);
  }

  return (
    <Container>
      <Tabs.Root defaultValue="application" onValueChange={onTabSwitched}>
        <Tabs.List>
          <Tabs.Trigger value="application">
            <Icon size="lg" color="gray.500">
              <HiIdentification />
              Application
            </Icon>
          </Tabs.Trigger>
          <Tabs.Trigger value="company">
            <Icon size="lg" color="gray.500">
              <HiOfficeBuilding />
              Company
            </Icon>
          </Tabs.Trigger>
          <Tabs.Trigger value="job">
            <Icon size="lg" color="gray.500">
              <HiBriefcase />
              Job
            </Icon>
          </Tabs.Trigger>
        </Tabs.List>
        <Tabs.Content value="application">
          <form onSubmit={onApplicationSubmit}>
            <Fieldset.Root>
              <Fieldset.Legend>Update Application Status</Fieldset.Legend>

              <Fieldset.Content>
                <Stack gap="4" align="flex-start" maxW="sm">
                  <Input
                    hidden
                    value={userTabType}
                    {...register("updateData.applicationData.id")}
                  ></Input>

                  <Field.Root
                    invalid={!!errors.updateData?.applicationData?.status}
                  >
                    <Field.Label>Status</Field.Label>
                    <NativeSelect.Root size="sm" width="240px">
                      <NativeSelect.Field
                        placeholder="Select option"
                        {...register("updateData.applicationData.status")}
                      >
                        <For each={applicationStatusTypes}>
                          {(item, index) => (
                            <option key={index} value={item.value}>{item.label}</option>
                          )}
                        </For>
                      </NativeSelect.Field>
                      <NativeSelect.Indicator />
                    </NativeSelect.Root>
                    <Field.ErrorText>
                      {errors.updateData?.applicationData?.status?.message}
                    </Field.ErrorText>
                  </Field.Root>

                  <Button type="submit">Submit</Button>
                </Stack>
              </Fieldset.Content>
            </Fieldset.Root>
          </form>
        </Tabs.Content>
        <Tabs.Content value="company">
          <form onSubmit={onCompanySubmit}>
            <Fieldset.Root>
              <Fieldset.Legend>Update Company Details</Fieldset.Legend>

              <Fieldset.Content>
                <Stack gap="4" align="flex-start" maxW="sm">
                  <Input
                    hidden
                    value={userTabType}
                    {...register("updateData.companyData.id")}
                  ></Input>

                  <Field.Root
                    invalid={!!errors.updateData?.companyData?.companyName}
                  >
                    <Field.Label>Company Name</Field.Label>
                    <Input
                      {...register("updateData.companyData.companyName")}
                    />
                    <Field.ErrorText>
                      {errors.updateData?.companyData?.companyName?.message}
                    </Field.ErrorText>
                  </Field.Root>
                  <Field.Root
                    invalid={!!errors.updateData?.companyData?.description}
                  >
                    <Field.Label>Description</Field.Label>
                    <Input
                      {...register("updateData.companyData.description")}
                    />
                    <Field.ErrorText>
                      {errors.updateData?.companyData?.description?.message}
                    </Field.ErrorText>
                  </Field.Root>

                  <Field.Root
                    invalid={!!errors.updateData?.companyData?.website}
                  >
                    <Field.Label>Website</Field.Label>
                    <Input {...register("updateData.companyData.website")} />
                    <Field.ErrorText>
                      {errors.updateData?.companyData?.website?.message}
                    </Field.ErrorText>
                  </Field.Root>

                  <Button type="submit">Submit</Button>
                </Stack>
              </Fieldset.Content>
            </Fieldset.Root>
          </form>
        </Tabs.Content>
        <Tabs.Content value="job">
          <form onSubmit={onJobSubmit}>
            <Fieldset.Root>
              <Fieldset.Legend>Update Job Details</Fieldset.Legend>
              <Fieldset.Content>
                <Stack gap="4" align="flex-start" maxW="sm">
                  <Input
                    hidden
                    value={userTabType}
                    {...register("updateData.jobData.id")}
                  ></Input>

                  <Field.Root invalid={!!errors.updateData?.jobData?.title}>
                    <Field.Label>Job Title</Field.Label>
                    <Input {...register("updateData.jobData.title")} />
                    <Field.ErrorText>
                      {errors.updateData?.jobData?.title?.message}
                    </Field.ErrorText>
                  </Field.Root>
                  <Field.Root
                    invalid={!!errors.updateData?.jobData?.description}
                  >
                    <Field.Label>Description</Field.Label>
                    <Input {...register("updateData.jobData.description")} />
                    <Field.ErrorText>
                      {errors.updateData?.jobData?.description?.message}
                    </Field.ErrorText>
                  </Field.Root>

                  <Field.Root
                    invalid={!!errors.updateData?.jobData?.requirements}
                  >
                    <TagsInput.Root delimiter={SPLIT_REGEX}>
                      <TagsInput.Label>Requirements</TagsInput.Label>
                      <TagsInput.Control>
                        <TagsInput.Items />

                        <TagsInput.Input placeholder="Type and use ; or Enter to add requirements..." />
                        <TagsInput.ClearTrigger />
                      </TagsInput.Control>

                      <TagsInput.HiddenInput />
                    </TagsInput.Root>
                    <Field.HelperText>Add job requirements</Field.HelperText>
                    <Field.ErrorText>
                      {errors.updateData?.jobData?.requirements?.message}
                    </Field.ErrorText>
                  </Field.Root>

                  <Button type="submit">Submit</Button>
                </Stack>
              </Fieldset.Content>
            </Fieldset.Root>
          </form>
        </Tabs.Content>
      </Tabs.Root>
    </Container>
  );
};

export default UpdatePage;
