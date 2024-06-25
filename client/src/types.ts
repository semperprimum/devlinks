import type { Component } from "vue";

export interface UserInfo {
  id: number;
  email: string;
  first_name: SqlStringOrNull;
  last_name: SqlStringOrNull;
  pic_path: SqlStringOrNull;
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
