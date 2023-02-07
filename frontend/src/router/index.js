import { createRouter, createWebHistory } from "vue-router";
import SubmitForm from "../views/SubmitForm";
import MessagesHistory from "../views/MessagesHistory";
import HomePage from "../views/Home";

const routes = [
    {
        path: "/",
        name: "HomePage",
        component: HomePage
    },
    {
        path: "/submit",
        name: "SubmitForm",
        component: SubmitForm
    },
    {
        path: "/messages",
        name: "MessagesHistory",
        component: MessagesHistory
    }
];

const router = createRouter({
    history: createWebHistory(process.env.BASE_URL),
    routes
});

export default router;
