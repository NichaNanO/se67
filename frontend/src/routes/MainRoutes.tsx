import { lazy } from "react";

import React from "react";

import { RouteObject } from "react-router-dom";

import MinimalLayout from "../layout/MinimalLayout";

import Loadable from "../components/third-patry/Loadable";
import ContractPage from "../pages/contract";


const MainPages = Loadable(lazy(() => import("../pages/authentication/Login")));

const Registerages = Loadable(

  lazy(() => import("../pages/authentication/Register"))

);


const MainRoutes = (): RouteObject => {
  return {
    path: "/",
    element: <MinimalLayout />,
    children: [
      {
        path: "/",
        element: <MainPages />,
      },
      {
        path: "/signup",
        element: <Registerages />,
      },
      {
        path: "/contract", // เพิ่มเส้นทางใหม่ใน MainRoutes
        element: <ContractPage />, // สามารถเข้าถึงได้โดยไม่ต้องล็อกอิน
      },
      {
        path: "*",
        element: <MainPages />,
      },
    ],
  };
};


export default MainRoutes;