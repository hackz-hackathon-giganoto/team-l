import { useRouter } from "next/router";
import React from "react";
import { BookPage } from "~/components/domain/bookPage";

const BookDetail = () => {
  const router = useRouter();
  return(
    <div>
      <BookPage/>
      book:{router.query.bookId}
    </div>
  );
};

export default BookDetail;