import { useState } from "react";
import {
  Button,
  Stack,
  Field,
  Input,
  Tabs,
  Icon,
  Heading,
  RadioGroup,
  HStack,
  Container,
} from "@chakra-ui/react";
import { useForm } from "react-hook-form";
import { HiUser, HiOfficeBuilding, HiDotsHorizontal } from "react-icons/hi";

type userTabTypedData = {
  applicantData: {
    username: string;
    name: string;
  };
  companyData: {
    username: string;
    name: string;
    companyName: string;
    website: string;
    description: string;
  };
};

interface FormValues {
  userData: userTabTypedData;
  userType: string;
}

const userTabTypes = [
  { label: "Applicant", value: "applicant" },
  { label: "Company", value: "company" },
];

const HomePage = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
    reset,
  } = useForm<FormValues>({
    defaultValues: {
      userData: {
        applicantData: {
          username: "",
          name: "",
        },
        companyData: {
          username: "",
          name: "",
          companyName: "",
          website: "",
          description: "",
        },
      },
    },
  });
  const onSubmit = handleSubmit((data) => {
    console.log("submitted data: ", data);
  });
  const [userTabType, setUserTabType] = useState<string>("applicant");
  const onTabSwitched = (details: Tabs.TabsValueChangeDetails) => {
    // TODO: may be related to controller, do not reset if partially filled
    reset();
    setUserTabType(details.value);
  }

  return (
    <Container>
      <Tabs.Root defaultValue="applicant" onValueChange={onTabSwitched}>
        <Tabs.List>
          <Tabs.Trigger value="applicant">
            <Icon size="lg" color="gray.500">
              <HiUser />
              Applicant
            </Icon>
          </Tabs.Trigger>
          <Tabs.Trigger value="company">
            <Icon size="lg" color="gray.500">
              <HiOfficeBuilding />
              Company
            </Icon>
          </Tabs.Trigger>
          <Tabs.Trigger value="random">
            <Icon size="lg" color="gray.500">
              <HiDotsHorizontal />
              Random
            </Icon>
          </Tabs.Trigger>
        </Tabs.List>
        <Tabs.Content value="applicant">
          <Heading size="xl">Create Applicant</Heading>

          <form onSubmit={onSubmit}>
            <Stack gap="4" align="flex-start" maxW="sm">
              <Input
                hidden
                value={userTabType}
                {...register("userType")}
              ></Input>

              <Field.Root invalid={!!errors.userData?.applicantData?.username}>
                <Field.Label>Username</Field.Label>
                <Input {...register("userData.applicantData.username")} />
                <Field.ErrorText>
                  {errors.userData?.applicantData?.username?.message}
                </Field.ErrorText>
              </Field.Root>

              <Field.Root invalid={!!errors.userData?.applicantData?.name}>
                <Field.Label>Name</Field.Label>
                <Input {...register("userData.applicantData.name")} />
                <Field.ErrorText>
                  {errors.userData?.applicantData?.name?.message}
                </Field.ErrorText>
              </Field.Root>

              <Button type="submit">Submit</Button>
            </Stack>
          </form>
        </Tabs.Content>
        <Tabs.Content value="company">
          <Heading size="xl">Create Company</Heading>

          <form onSubmit={onSubmit}>
            <Stack gap="4" align="flex-start" maxW="sm">
              <Input
                hidden
                value={userTabType}
                {...register("userType")}
              ></Input>

              <Field.Root invalid={!!errors.userData?.companyData?.username}>
                <Field.Label>Owner Username</Field.Label>
                <Input {...register("userData.companyData.username")} />
                <Field.ErrorText>
                  {errors.userData?.companyData?.username?.message}
                </Field.ErrorText>
              </Field.Root>

              <Field.Root invalid={!!errors.userData?.companyData?.name}>
                <Field.Label>Owner Name</Field.Label>
                <Input {...register("userData.companyData.name")} />
                <Field.ErrorText>
                  {errors.userData?.companyData?.name?.message}
                </Field.ErrorText>
              </Field.Root>

              <Field.Root invalid={!!errors.userData?.companyData?.companyName}>
                <Field.Label>Company Name</Field.Label>
                <Input {...register("userData.companyData.companyName")} />
                <Field.ErrorText>
                  {errors.userData?.companyData?.companyName?.message}
                </Field.ErrorText>
              </Field.Root>
              <Field.Root invalid={!!errors.userData?.companyData?.description}>
                <Field.Label>Description</Field.Label>
                <Input {...register("userData.companyData.description")} />
                <Field.ErrorText>
                  {errors.userData?.companyData?.description?.message}
                </Field.ErrorText>
              </Field.Root>

              <Field.Root invalid={!!errors.userData?.companyData?.website}>
                <Field.Label>Website</Field.Label>
                <Input {...register("userData.companyData.website")} />
                <Field.ErrorText>
                  {errors.userData?.companyData?.website?.message}
                </Field.ErrorText>
              </Field.Root>

              <Button type="submit">Submit</Button>
            </Stack>
          </form>
        </Tabs.Content>
        <Tabs.Content value="random">
          <Stack align="flex-start" maxW="sm">
            <Heading size="xl">Create random user (WIP)</Heading>
            <RadioGroup.Root defaultValue="applicant">
              <HStack gap="6">
                {userTabTypes.map((item) => (
                  <RadioGroup.Item key={item.value} value={item.value}>
                    <RadioGroup.ItemHiddenInput />
                    <RadioGroup.ItemIndicator />
                    <RadioGroup.ItemText>{item.label}</RadioGroup.ItemText>
                  </RadioGroup.Item>
                ))}
              </HStack>
            </RadioGroup.Root>
            {/* TODO */}
            <Button disabled>Generate</Button>
          </Stack>
        </Tabs.Content>
      </Tabs.Root>
    </Container>
  );
};

export default HomePage;
