import { useRouter } from "next/router";
import React from "react";

const DmDetail = () => {
  const router = useRouter();
  return(
    <div>user:{router.query.dmId}</div>
  )
}

export default DmDetail;
