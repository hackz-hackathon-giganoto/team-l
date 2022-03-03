import { useRouter } from "next/router"
import React from "react"
import { UserPage } from "~/components/domain/userPage"

const User = () => {
  const router = useRouter()
  return(
    <div>
      <UserPage/>
      userName:{router.query.userName}
    </div>
  )
}

export default User;