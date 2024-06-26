import { PlatformColors } from "@/types";
import type { Component } from "vue";
import GitHub from "@/components/icons/GitHub.vue";
import FrontendMentor from "@/components/icons/FrontendMentor.vue";
import Twitter from "@/components/icons/Twitter.vue";
import Linkedin from "@/components/icons/Linkedin.vue";
import YouTube from "@/components/icons/YouTube.vue";
import Facebook from "@/components/icons/Facebook.vue";
import Twitch from "@/components/icons/Twitch.vue";
import Devto from "@/components/icons/Devto.vue";
import Codewars from "@/components/icons/Codewars.vue";
import Codepen from "@/components/icons/Codepen.vue";
import FreeCodeCamp from "@/components/icons/freeCodeCamp.vue";
import GitLab from "@/components/icons/GitLab.vue";
import Hashnode from "@/components/icons/Hashnode.vue";
import StackOverflow from "@/components/icons/StackOverflow.vue";

// prettier-ignore
const platformProperties: { [key: string]: { color: string; icon: Component; name: string } } = {
  github: { color: PlatformColors.github, icon: GitHub, name: "GitHub" },
  frontendmentor: { color: PlatformColors.frontendmentor, icon: FrontendMentor, name: "Frontend Mentor" },
  twitter: { color: PlatformColors.twitter, icon: Twitter, name: "Twitter" },
  linkedin: { color: PlatformColors.linkedin, icon: Linkedin, name: "LinkedIn" },
  youtube: { color: PlatformColors.youtube, icon: YouTube, name: "YouTube" },
  facebook: { color: PlatformColors.facebook, icon: Facebook, name: "Facebook" },
  twitch: { color: PlatformColors.twitch, icon: Twitch, name: "Twitch" },
  devto: { color: PlatformColors.devto, icon: Devto, name: "Dev.to" },
  codewars: { color: PlatformColors.codewars, icon: Codewars, name: "Codewars" },
  codepen: { color: PlatformColors.codepen, icon: Codepen, name: "Codepen" },
  freecodecamp: { color: PlatformColors.freecodecamp, icon: FreeCodeCamp, name: "freeCodeCamp" },
  gitlab: { color: PlatformColors.gitlab, icon: GitLab, name: "GitLab" },
  hashnode: { color: PlatformColors.hashnode, icon: Hashnode, name: "Hashnode" },
  stackoverflow: { color: PlatformColors.stackoverflow, icon: StackOverflow, name: "Stack Overflow" },
};

export const getPlatformProperties = (
  name: string,
): { color: string; icon: Component; name: string } => {
  return (
    platformProperties[name] || {
      color: PlatformColors.github,
      icon: GitHub,
      name: "GitHub",
    }
  );
};
