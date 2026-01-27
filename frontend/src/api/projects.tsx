import axios from "axios";

export type Response = {
  data: SuccessfulData & ErrorData;
  status: number;
};

export type SuccessfulData = {
  code: number;
  data: Project[];
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

export async function getProjects(): Promise<Response> {
  return await axios.get(`${import.meta.env.VITE_API_URL}/project`);
}
