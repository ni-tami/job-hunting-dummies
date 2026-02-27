import { useState } from "react";
import {
  Stack,
  Button,
  Container,
  Box,
  VStack,
  For,
  Text,
  Link,
  Tabs,
  Collapsible,
  List,
  TabsContent,
  Editable,
  IconButton,
  Field,
  NativeSelect,
} from "@chakra-ui/react";
import { useForm } from "react-hook-form";
import {
  LuBox,
  LuPencilLine,
  LuX,
  LuCheck,
} from "react-icons/lu";
import {
  dummyApplications,
  dummyCompanies,
  dummyJobs,
} from "@/utils/dummies";
import { Company, Application, Job } from "@/types";
import { applicationStatusTypes } from "@/constants";


const TabComponentMapping = [
  {
    tab: {
      value: "application",
      display: "Application",
    },
    content: (
      <Tabs.Content value="application">
        Manage Applications
        <For
          each={dummyApplications}
          fallback={
            <VStack textAlign="center" fontWeight="medium">
              <LuBox />
              No items to show
            </VStack>
          }
        >
          {(appl: Application, applIdx: number) => (
            <Box borderWidth="1px" key={applIdx} p="4">
              <Text fontWeight="bold">{appl.id}</Text>
              <Editable.Root defaultValue={appl.status} activationMode="none">
                <NativeSelect.Root size="sm" width="240px">
                  <NativeSelect.Field
                    defaultValue={appl.status}
                  >
                    <For each={applicationStatusTypes}>
                      {(item, index) => (
                        <option key={index} value={item.value}>
                          {item.label}
                        </option>
                      )}
                    </For>
                  </NativeSelect.Field>
                  <NativeSelect.Indicator />
                </NativeSelect.Root>
                <Editable.Control>
                  <Editable.EditTrigger asChild>
                    <IconButton variant="ghost" size="xs">
                      <LuPencilLine />
                    </IconButton>
                  </Editable.EditTrigger>
                  <Editable.CancelTrigger asChild>
                    <IconButton variant="outline" size="xs">
                      <LuX />
                    </IconButton>
                  </Editable.CancelTrigger>
                  <Editable.SubmitTrigger asChild>
                    <IconButton variant="outline" size="xs">
                      <LuCheck />
                    </IconButton>
                  </Editable.SubmitTrigger>
                </Editable.Control>
              </Editable.Root>
            </Box>
          )}
        </For>
      </Tabs.Content>
    ),
  },
  {
    tab: {
      value: "company",
      display: "Company",
    },
    content: (
      <Tabs.Content value="company">
        Manage Companies
        <For
          each={dummyCompanies}
          fallback={
            <VStack textAlign="center" fontWeight="medium">
              <LuBox />
              No items to show
            </VStack>
          }
        >
          {(company: Company, companyIdx: number) => (
            <Box borderWidth="1px" key={companyIdx} p="4">
              <Text fontWeight="bold">
                Company No. {company.id} - {company.companyName}
              </Text>
              <Text color="fg.muted">{company.description}</Text>
              <Link variant="underline" href={company.website} color="blue.fg">
                Visit company website
              </Link>
            </Box>
          )}
        </For>
      </Tabs.Content>
    ),
  },
  {
    tab: {
      value: "job",
      display: "Jobs",
    },
    content: (
      <Tabs.Content value="job">
        Manage Jobs
        <For
          each={dummyJobs}
          fallback={
            <VStack textAlign="center" fontWeight="medium">
              <LuBox />
              No items to show
            </VStack>
          }
        >
          {(job: Job, jobIdx: number) => (
            <Box borderWidth="1px" key={jobIdx} p="4">
              <Text fontWeight="bold">
                JobID: {job.id} - {job.title}
              </Text>
              <Text color="fg.muted">{job.description}</Text>
              <Collapsible.Root>
                <Collapsible.Trigger paddingY="3">
                  Requirements
                </Collapsible.Trigger>
                <Collapsible.Content>
                  <List.Root>
                    <For each={job.requirements}>
                      {(req, reqIdx) => (
                        <List.Item key={reqIdx}>{req}</List.Item>
                      )}
                    </For>
                  </List.Root>
                </Collapsible.Content>
              </Collapsible.Root>
            </Box>
          )}
        </For>
      </Tabs.Content>
    ),
  },
];

const CompanyDashboardPage = () => {
  return (
    <Container>
      <Tabs.Root defaultValue="application" variant="plain">
        <Tabs.List bg="bg.muted" rounded="l3" p="1">
          <For each={TabComponentMapping}>
            {(tabComponent, _) => (
              <Tabs.Trigger value={tabComponent.tab.value}>
                {tabComponent.tab.display}
              </Tabs.Trigger>
            )}
          </For>
          <Tabs.Indicator rounded="l2" />
        </Tabs.List>
        <For each={TabComponentMapping}>
          {(tabComponent) => <>{tabComponent.content}</>}
        </For>
      </Tabs.Root>
    </Container>
  );
};

export default CompanyDashboardPage;
