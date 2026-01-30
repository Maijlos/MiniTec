import axios from "axios";

type Response = {data: SuccessfulData | ErrorData};

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
  const data: Response = await axios.get(`${import.meta.env.VITE_API_URL}/project`);
  return data.data;
}

export async function getProjectHealth(id: number) {
  const data = await axios.get(
    `${import.meta.env.VITE_API_URL}/project/health/${id}`,
  );
  return data.data;
}

export async function deleteProject(id: number) {
  const data = await axios.delete(
    `${import.meta.env.VITE_API_URL}/project/${id}`,
  );
  return data.data;
}
