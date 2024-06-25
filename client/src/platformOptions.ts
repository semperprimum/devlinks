import Codepen from "@/components/icons/Codepen.vue";
import Codewars from "@/components/icons/Codewars.vue";
import Devto from "@/components/icons/Devto.vue";
import Facebook from "@/components/icons/Facebook.vue";
import FrontendMentor from "@/components/icons/FrontendMentor.vue";
import GitHub from "@/components/icons/GitHub.vue";
import GitLab from "@/components/icons/GitLab.vue";
import Hashnode from "@/components/icons/Hashnode.vue";
import Linkedin from "@/components/icons/Linkedin.vue";
import StackOverflow from "@/components/icons/StackOverflow.vue";
import Twitch from "@/components/icons/Twitch.vue";
import Twitter from "@/components/icons/Twitter.vue";
import YouTube from "@/components/icons/YouTube.vue";
import FreeCodeCamp from "@/components/icons/freeCodeCamp.vue";
import type { SelectOption } from "@/types";

export const platformOptions: SelectOption[] = [
  { value: "github", label: "GitHub", icon: GitHub },
  { value: "frontendmentor", label: "Frontend Mentor", icon: FrontendMentor },
  { value: "twitter", label: "Twitter", icon: Twitter },
  { value: "linkedin", label: "LinkedIn", icon: Linkedin },
  { value: "youtube", label: "YouTube", icon: YouTube },
  { value: "facebook", label: "Facebook", icon: Facebook },
  { value: "twitch", label: "Twitch", icon: Twitch },
  { value: "devto", label: "Dev.to", icon: Devto },
  { value: "codewars", label: "Codewars", icon: Codewars },
  { value: "codepen", label: "Codepen", icon: Codepen },
  { value: "freecodecamp", label: "freeCodeCamp", icon: FreeCodeCamp },
  { value: "gitlab", label: "GitLab", icon: GitLab },
  { value: "hashnode", label: "Hashnode", icon: Hashnode },
  { value: "stackoverflow", label: "Stack Overflow", icon: StackOverflow },
];
