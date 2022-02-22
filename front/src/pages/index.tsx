import type { NextPage } from "next";
import useSWR from "swr";

const fetcher = (url: string) => fetch(url).then((res) => res.json());

const Home: NextPage = () => {
  const { data, error } = useSWR("http://localhost:8080/", fetcher);

  if (error) return <div>failed to load</div>;
  if (!data) return <div>loading...</div>;

  return <div>{data.msg}</div>;
};

export default Home;
