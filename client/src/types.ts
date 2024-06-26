import type { Component } from "vue";

export interface UserInfo {
  id: number;
  email: string;
  display_email: SqlStringOrNull;
  first_name: SqlStringOrNull;
  last_name: SqlStringOrNull;
  pic_path: SqlStringOrNull;
}
export interface UserLinks {
  id: number;
  url: string;
  platform: string;
  order_index: number;
}

export interface GetUserLinksResponse {
  user: {
    id: number;
    display_email: SqlStringOrNull;
    first_name: SqlStringOrNull;
    last_name: SqlStringOrNull;
    pic_path: SqlStringOrNull;
  };
  links: UserLinks[];
}

type SqlStringOrNull = {
  String: string;
  Valid: boolean;
};

export interface SelectOption {
  value: string | number;
  label: string;
  icon?: Component;
}

export enum PlatformColors {
  github = "#1A1A1A",
  devto = "#333333",
  frontendmentor = "#FFFFFF",
  codepen = "#AE63E4",
  codewars = "#8A1A50",
  twitter = "#43B7E9",
  freecodecamp = "#302267",
  linkedin = "#2D68FF",
  gitlab = "#EB4925",
  youtube = "#EE3939",
  hashnode = "#0330D1",
  facebook = "#2442AC",
  stackoverflow = "#EC7100",
  twitch = "#EE3FC8",
}
