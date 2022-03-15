import { Button } from "@nextui-org/react";
import React from "react";
import { giganotoApi } from "~/api/api";

export const UserPage = () => {
  const DUMMY_DATA = "c8o18gu49b3hs0ma51cg";
  const userId = DUMMY_DATA;

  const userData = async() => {
    const user = await giganotoApi.fetchUserDetail(userId);
    console.log(user);
  };

  return(
    <div>
      <Button onClick={userData}>button</Button>
    </div>
  );
};