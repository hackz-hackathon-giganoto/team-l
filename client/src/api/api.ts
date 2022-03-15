import axios from "axios";
const BASE_URL_LOCAL = "http://954f-49-104-48-57.ngrok.io";

const giganotoAxios = axios.create({
  baseURL: BASE_URL_LOCAL
});

export type ArticleBody = {
  userName: string,
  introduction: string,
  image: string,
  twitter: string,
  github: string
};

export type UserRegistBody = {
  userName: string,
  introduction: string,
  image: string,
  twitter: string,
  github: string
}

export const giganotoApi = {
  fetchUserDetail: async (userId:string) => {
    const data = await giganotoAxios.get(`/user/${userId}`);
    return data;
  },
  fetchArticle: async (articleId: string) => {
    const data = await giganotoAxios.get(`/books/${articleId}`);
    return data;
  },
  postArticle: async (userId: string,body: ArticleBody) => {
    const log = await giganotoAxios.post(`/books/${userId}`, body);
    return log;
  },
  regist: async (body: UserRegistBody) => {
    const log = await giganotoAxios.post(`/user`, body);
    return log;
  }
};
