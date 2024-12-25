import { lazy } from "react";
import { RouteObject } from "react-router-dom";
import Loadable from "../components/third-patry/Loadable";
import FullLayout from "../layout/FullLayout";

// โหลดหน้าใหม่ contractPage
const MainPages = Loadable(lazy(() => import("../pages/authentication/Login")));
const Dashboard = Loadable(lazy(() => import("../pages/dashboard")));
const Customer = Loadable(lazy(() => import("../pages/customer")));
const CreateCustomer = Loadable(lazy(() => import("../pages/customer/create")));
const EditCustomer = Loadable(lazy(() => import("../pages/customer/edit")));
const ContractPage = Loadable(lazy(() => import("../pages/contract"))); // เพิ่มหน้า contractPage

const AdminRoutes = (isLoggedIn: boolean): RouteObject => {
  return {
    path: "/",
    element: isLoggedIn ? <FullLayout /> : <MainPages />,
    children: [
      {
        path: "/",
        element: <Dashboard />,
      },
      {
        path: "/customer",
        children: [
          {
            path: "/customer",
            element: <Customer />,
          },
          {
            path: "/customer/create",
            element: <CreateCustomer />,
          },
          {
            path: "/customer/edit/:id",
            element: <EditCustomer />,
          },
        ],
      },
      // เพิ่มเส้นทางใหม่สำหรับ contractPage
      {
        path: "/contract",
        element: <ContractPage />, // ระบุ Component ที่ต้องการให้แสดง
      },
    ],
  };
};

export default AdminRoutes;
