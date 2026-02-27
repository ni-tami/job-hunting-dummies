import { Company, Application, Job } from "@/types";

export const dummyCompanies: Company[] = [
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

export const dummyApplications: Application[] = [
  {
    id: 34820813,
    status: "APPLIED",
  },
  {
    id: 10239035,
    status: "ACCEPTED",
  },
];

export const dummyJobs: Job[] = [
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
