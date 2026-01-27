import { useQuery } from "@tanstack/react-query";
import { projectKeys } from "../api/keys.tsx";
import { getProjects } from "../api/projects.tsx";
import { DropdownMenu } from "./DropdownMenu.tsx";
import { useState } from "react";

export function FetchProjects() {
  const { data, isError, isLoading } = useQuery({
    queryKey: projectKeys.allProjects,
    queryFn: getProjects,
  });

  const [chosedProject, setChosedProject] = useState<number | null>(null);

  if (isError) {
    return <div>Something went wrong!</div>;
  }

  if (isLoading) {
    return <div>Loading...</div>;
  }

  const projectId = chosedProject ?? data?.data.data[0].id;
  console.log(projectId);


  return (
    <>
      <DropdownMenu projects={data?.data} setChosedProject={setChosedProject} />
    </>
  );
}
