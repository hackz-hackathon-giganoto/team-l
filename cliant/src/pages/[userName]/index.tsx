import { useRouter } from "next/router"
import React from "react"

const User = () => {
  const router = useRouter()
  return(
    <div>userName:{router.query.userName}</div>
  )
}

export default User;