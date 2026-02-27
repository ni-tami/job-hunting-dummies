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
} from "@chakra-ui/react";
import { useForm } from "react-hook-form";
import {
  LuBox,
  LuUser
} from "react-icons/lu";
import {
  UpdateCompanyMutationTmpl,
  UpdateApplicationMutationTmpl,
  UpdateJobMutationTmpl,
} from "@/utils/graphql";

type Application = {
  id: number;
  status: string;
};

type Company = {
  id: number;
  companyName: string;
  website: string;
  description: string;
};

type Job = {
  id: number;
  title: string;
  description: string;
  requirements: string[];
};

const dummyCompanies: Company[] = [
  {
    id: 123,
    companyName: "hehe",
    description: "hehe is a well-known good company",
    website: "https://www.hehe.com",
  },
  {
    id: 456,
    companyName: "hoho",
    description: "hoho is a healthy company",
    website: "https://www.hoho.com",
  },
];

const dummyApplications: Application[] = [
  {
    id: 34820813,
    status: "APPLIED",
  },
  {
    id: 10239035,
    status: "ACCEPTED",
  },
];

const dummyJobs: Job[] = [
  {
    id: 12315,
    title: "Data Wrangling",
    description: "Help get insights from given dataset",
    requirements: ["Clean data", "Visualize data", "Gain insights"],
  },
  {
    id: 24593,
    title: "Draw Fishies",
    description: "Paint smol fishies for room decor",
    requirements: ["Can paint with water colors"],
  },
];

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
              <Text fontWeight="bold">{appl.status}</Text>
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
          {(job, jobIdx) => (
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
      <Tabs.Root defaultValue="members" variant="plain">
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
