import { useRouter } from "next/router";
import React from "react"

const BookDetail = () => {
  const router = useRouter();
  
  return(
    <div>
      book:{router.query.bookId}
    </div>
  )
}

export default BookDetail