import { useQuery } from "@tanstack/react-query";
import { projectKeys } from "../api/keys.tsx";
import { getProjects } from "../api/projects.tsx";
import { DropdownMenu } from "./DropdownMenu.tsx";
import { Table } from "./Table.tsx";
import { useState } from "react";

export function FetchProjects() {
  const { data, isError, isLoading } = useQuery({
    queryKey: projectKeys.allProjects,
    queryFn:  getProjects,
  });

  const [chosedProject, setChosedProject] = useState<number | null>(null);

  if (isError) {
    return <div className="flex justify-center text-4xl">Something went wrong!</div>;
  }

  if (isLoading) {
    return <div className="flex justify-center text-4xl">Loading...</div>;
  }

  const projectId = chosedProject ?? (data && 'data' in data && Array.isArray(data.data) ? data.data[0].id : null)

  return (
    <>
      {(data && 'data' in data && Array.isArray(data.data)) && <DropdownMenu projects={data} setChosedProject={setChosedProject} />}
      {projectId && <Table id={projectId} />}
    </>
  );
}
