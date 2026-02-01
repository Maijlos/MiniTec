import axios from "./instance";

type Response = { data: SuccessfulData | ErrorData };

type Stations = {
  [name: string]: [
    {
      finalStatus: string;
      startDT: string;
      endDT: string;
    },
  ];
};

export type SuccessfulData = {
  code: number;
  data: Project[] | Stations;
  shortMessage: "success";
};

export type ErrorData = {
  code: number;
  message: string;
  shortMessage: string;
};

export type Project = {
  code: string;
  id: number;
  name: string;
};

export async function getProjects(): Promise<SuccessfulData | ErrorData> {
  const data: Response = await axios.get(`/project`);
  return data.data;
}

export async function getProject(id: number) {
  const data = await axios.get(`/project/${id}`);
  return data.data;
}

export async function createProject(
  name: string,
  code: string,
): Promise<SuccessfulData | ErrorData> {
  const data: Response = await axios.post(`/project`, { name, code });
  return data.data;
}

export function updateProject(id: number, name: string, code: string) {
  return axios.put(`/project/${id}`, { name, code });
}

export async function getProjectHealth(id: number) {
  const data = await axios.get(`/project/health/${id}`);
  return data.data;
}

export async function deleteProject(id: number) {
  const data = await axios.delete(`/project/${id}`);
  return data.data;
}
