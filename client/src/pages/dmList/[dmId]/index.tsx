import { useRouter } from "next/router";
import React from "react";
import { DmPage } from "~/components/domain/dmPage";

const Dm = () => {
  const router = useRouter();
  return(
    <div>
      <DmPage/>
      user:{router.query.dmId}
    </div>
  );
};

export default Dm;
