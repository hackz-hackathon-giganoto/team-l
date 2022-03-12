import { AppProps } from "next/app";
import { NextUIProvider } from "@nextui-org/react";
import { SessionProvider } from "next-auth/react";

const MyApp = ({ Component, pageProps }: AppProps) => {
  return (
    <NextUIProvider>
      <SessionProvider session={pageProps.session}>
        <Component {...pageProps} />
      </SessionProvider>
    </NextUIProvider>
  );
};

export default MyApp;
