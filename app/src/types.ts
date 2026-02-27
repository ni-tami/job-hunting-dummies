export type Application = {
  id: number;
  status: string;
};

export type Company = {
  id: number;
  companyName: string;
  website: string;
  description: string;
};

export type Job = {
  id: number;
  title: string;
  description: string;
  requirements: string[];
};
