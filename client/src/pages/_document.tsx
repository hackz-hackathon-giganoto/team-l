import { Html, Head, Main, NextScript, DocumentContext, DocumentInitialProps } from 'next/document';
import { CssBaseline } from '@nextui-org/react';
import { NextPage } from 'next';

const MyDocument: NextPage = () => {
  return(
    <Html>
      <Head>{CssBaseline.flush()}</Head>
      <body>
        <Main />
        <NextScript />
      </body>
    </Html>
  );
};

MyDocument.getInitialProps = async (ctx:DocumentContext): Promise<DocumentInitialProps> => {
  const initialProps = await ctx.defaultGetInitialProps(ctx);
  console.log("getInitialProps");
    return {
      ...initialProps,
      styles: <>{initialProps.styles}</>
    };
};

export default MyDocument;
