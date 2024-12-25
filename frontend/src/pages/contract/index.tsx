import React, { useEffect, useState, useRef } from "react";
import axios from "axios";
import { Form, Input, Button, DatePicker, Select, Upload, message } from "antd";
import { UploadOutlined } from "@ant-design/icons";
import dayjs from "dayjs";
import SignatureCanvas from "react-signature-canvas";

const ContractPage: React.FC = () => {
  const [form] = Form.useForm();
  const sigCanvas = useRef<SignatureCanvas>(null);

  const [members, setMembers] = useState<any[]>([]);
  const [rooms, setRooms] = useState<any[]>([]); // Ensure rooms is an array by default
  const [contractTypes, setContractTypes] = useState<any[]>([]);
  const [employeeId, setEmployeeId] = useState<number | null>(null);
  const [securityDeposit, setSecurityDeposit] = useState<number>(0);

  // Fetch members, rooms, and contract types from the API
  useEffect(() => {
    const fetchData = async () => {
      try {
        const [membersData, roomsData, contractTypesData] = await Promise.all([
          axios.get("/api/members"),
          axios.get("/api/rooms"),
          axios.get("/api/contracttypes"),
        ]);
  
        // ตรวจสอบว่าข้อมูล contractTypes เป็นอาร์เรย์หรือไม่
        const contractTypes = Array.isArray(contractTypesData.data) ? contractTypesData.data : [];
        setMembers(membersData.data);
        setRooms(Array.isArray(roomsData.data) ? roomsData.data : []);
        setContractTypes(contractTypes); // ตั้งค่า contractTypes หากเป็นอาร์เรย์
  
        // รับ employee_id จาก localStorage
        const storedEmployeeId = localStorage.getItem("employeeId");
        if (storedEmployeeId) {
          setEmployeeId(Number(storedEmployeeId));
        }
      } catch (error) {
        console.error("ไม่สามารถดึงข้อมูลได้", error);
      }
    };
    fetchData();
  }, []);
  

  // Calculate end date based on contract type and check-in date
  const calculateEndDate = (startDate: string, contractType: number) => {
    const durationMap: { [key: number]: number } = {
      3: 3,
      6: 6,
      9: 9,
    };
    const monthsToAdd = durationMap[contractType] || 0;
    return dayjs(startDate).add(monthsToAdd, "month").format("YYYY-MM-DD");
  };

  // Clear signature
  const clearSignature = () => {
    sigCanvas.current?.clear();
  };

  // Handle form submission
  const onFinish = async (values: any) => {
    try {
      const formattedDate = values.checkInDate
        ? dayjs(values.checkInDate).format("YYYY-MM-DD")
        : null;

      if (!formattedDate) {
        message.error("กรุณาเลือกวันที่เข้าพัก");
        return;
      }

      const endDate = calculateEndDate(formattedDate, values.contractTypeId);

      // Convert signature to base64
      const signature = sigCanvas.current?.toDataURL();

      const formData = {
        ...values,
        checkInDate: formattedDate,
        endDate: endDate,
        employeeId: employeeId, // Pass employee_id from localStorage
        securityDeposit: securityDeposit, // Pass the security deposit value
        signature: signature, // Pass signature in base64 format
      };

      // Send form data to API
      const response = await axios.post("http://localhost:8080/contracts", formData);

      if (response.status === 201) {
        message.success("บันทึกข้อมูลสำเร็จ!");
        form.resetFields();
        clearSignature(); // Clear the signature after successful save
      }
    } catch (error) {
      console.error("เกิดข้อผิดพลาด:", error);
      message.error("บันทึกข้อมูลล้มเหลว กรุณาลองใหม่อีกครั้ง");
    }
  };

  // Handle contract type change and update security deposit
  const handleContractTypeChange = (value: number) => {
    const contractType = contractTypes.find((type) => type.id === value);
    if (contractType) {
      setSecurityDeposit(contractType.securityDeposit); // Update security deposit based on contract type
    }
  };

  return (
    <div style={{ padding: "20px", background: "linear-gradient(to bottom, #e0eafc, #cfdef3)" }}>
      <Form form={form} layout="vertical" style={{ maxWidth: "900px", margin: "0 auto" }} onFinish={onFinish}>
        {/* Section 1: Personal Information */}
        <div style={{ display: "flex", justifyContent: "space-between", gap: "20px" }}>
          <div style={{ flex: 1 }}>
            <h3>ส่วนที่ 1: ข้อมูลส่วนตัว</h3>
            <Form.Item label="ชื่อ-สกุล / Full Name" name="fullName" rules={[{ required: true }]}>
              <Input placeholder="ชื่อ-สกุล" />
            </Form.Item>
            <Form.Item label="หมายเลขประจำตัวประชาชน" name="idCard" rules={[{ required: true }]}>
              <Input placeholder="1234567890123" />
            </Form.Item>
            <Form.Item label="หมายเลขโทรศัพท์ / Telephone" name="phone" rules={[{ required: true }]}>
              <Input placeholder="0619807818" />
            </Form.Item>
            <Form.Item label="ที่อยู่ / Address" name="address" rules={[{ required: true }]}>
              <Input placeholder="ที่อยู่..." />
            </Form.Item>
          </div>

          <div style={{ flex: 1 }}>
            <h3>ส่วนที่ 2: บัญชีผู้ใช้งาน</h3>
            <Form.Item label="รหัสผ่าน / Password" name="password" rules={[{ required: true }]}>
              <Input.Password placeholder="Password" />
            </Form.Item>
            <Form.Item label="ยืนยันรหัสผ่าน / Password Again" name="confirmPassword" rules={[{ required: true }]}>
              <Input.Password placeholder="Confirm Password" />
            </Form.Item>

            <h3 style={{ marginTop: "20px" }}>ส่วนที่ 3: ข้อมูลห้องพัก</h3>
            <Form.Item label="หมายเลขห้องพัก / Room Number" name="roomId" rules={[{ required: true }]}>
              <Select placeholder="เลือกหมายเลขห้อง">
                {Array.isArray(rooms) && rooms.map((room) => (
                  <Select.Option key={room.id} value={room.id}>
                    {room.roomNumber}
                  </Select.Option>
                ))}
              </Select>
            </Form.Item>
            <Form.Item label="วันเข้าพัก / Check-in Date" name="checkInDate" rules={[{ required: true }]}>
              <DatePicker style={{ width: "100%" }} />
            </Form.Item>
          </div>
        </div>

        {/* Section 4: Contract Type */}
        <Form.Item label="ประเภทสัญญา / Contract Type" name="contractTypeId" rules={[{ required: true }]}>
  <Select placeholder="เลือกประเภทสัญญา" onChange={handleContractTypeChange}>
    {Array.isArray(contractTypes) && contractTypes.map((contractType) => (
      <Select.Option key={contractType.id} value={contractType.id}>
        {contractType.contractName} (มัดจำ {contractType.securityDeposit} บาท)
      </Select.Option>
    ))}
  </Select>
</Form.Item>
        {/* Signature Section */}
        <div style={{ textAlign: "center", marginTop: "30px" }}>
          <h3>ลายเซ็น</h3>
          <SignatureCanvas ref={sigCanvas} penColor="black" canvasProps={{ width: 500, height: 200, className: "signature-canvas" }} />
          <div style={{ marginTop: "10px" }}>
            <Button onClick={clearSignature}>ล้างลายเซ็น</Button>
          </div>
        </div>

        {/* File Upload */}
        <div style={{ textAlign: "center", marginTop: "30px" }}>
          <h3>อัปโหลดรูปภาพ</h3>
          <Form.Item name="upload" style={{ display: "inline-block" }}>
            <Upload>
              <Button icon={<UploadOutlined />}>เลือกไฟล์</Button>
            </Upload>
          </Form.Item>
        </div>

        {/* Buttons */}
        <div style={{ textAlign: "center", marginTop: "20px" }}>
          <Form.Item>
            <Button type="primary" htmlType="submit" style={{ marginRight: "10px" }}>
              กดเพื่อสมัครสมาชิก / Register
            </Button>
            <Button htmlType="button" danger onClick={() => form.resetFields()}>
              ล้างข้อมูล / Reset
            </Button>
          </Form.Item>
        </div>
      </Form>
    </div>
  );
};

export default ContractPage;
