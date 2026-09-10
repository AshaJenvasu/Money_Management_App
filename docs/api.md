# 📡 API Documentation — Money Management App

> Base URL: `http://localhost:3000/api/v1`
> Auth: แนบ JWT ใน Header `Authorization: Bearer <token>` (ยกเว้น `/auth/*`)

---

## 📋 Template — คัดลอกไปใช้ทุก Endpoint ใหม่

```markdown
### [METHOD] /path

**Auth required:** Yes / No

**Description:** อธิบายสั้นๆ ว่า endpoint นี้ทำอะไร

**Request Body:**
\`\`\`json
{ }
\`\`\`

**Response 200:**
\`\`\`json
{ }
\`\`\`

**Errors:**
| Code | Reason |
|---|---|
| 400 | |
| 401 | |
| 404 | |
```

---

## ✅ Auth (ตัวอย่างที่กรอกแล้ว)

### POST /auth/register

**Auth required:** No

**Description:** สมัครสมาชิกใหม่

**Request Body:**
```json
{
  "email": "user@example.com",
  "username": "asha_dev",
  "password": "your_password"
}
```

**Response 201:**
```json
{
  "id": 1,
  "email": "user@example.com",
  "username": "asha_dev"
}
```

**Errors:**
| Code | Reason |
|---|---|
| 400 | email หรือ username ซ้ำ / ข้อมูลไม่ครบ |

---

### POST /auth/login

**Auth required:** No

**Description:** เข้าสู่ระบบ ด้วย email หรือ username ก็ได้

**Request Body:**
```json
{
  "identifier": "user@example.com",
  "password": "your_password"
}
```

**Response 200:**
```json
{
  "token": "eyJhbGciOi..."
}
```

**Errors:**
| Code | Reason |
|---|---|
| 401 | identifier หรือ password ไม่ถูกต้อง |

---

## ⚠️ กฎทั่วไปที่ใช้กับทุก Endpoint ด้านล่าง (ยกเว้น /auth/*)

1. ต้องแนบ JWT — ไม่งั้นตอบ `401 Unauthorized`
2. **ทุก Endpoint ที่มี `:id` หรือ `:walletId` ต้องเช็คว่าทรัพยากรนั้นเป็นของ user ที่ Login อยู่จริง** ไม่ใช่แค่เช็คว่า JWT ถูกต้อง (Authentication ≠ Authorization) — ถ้าไม่ใช่เจ้าของ ตอบ `403 Forbidden` ไม่ใช่ปล่อยให้เห็นข้อมูลคนอื่น

---

## 🗂️ Route Overview

| Method | Path | Purpose |
|---|---|---|
| GET | /wallets | ดูกระเป๋าทั้งหมด |
| POST | /wallets | สร้างกระเป๋าใหม่ |
| GET | /wallets/:id | ดูกระเป๋าเดียว |
| PUT | /wallets/:id | แก้ไขกระเป๋า |
| DELETE | /wallets/:id | ลบกระเป๋า |
| GET | /wallets/:id/balance | ยอดคงเหลือ (คำนวณสด) |
| GET | /wallets/:walletId/transactions | ดูรายการ |
| POST | /wallets/:walletId/transactions | บันทึกรายการใหม่ |
| GET | /wallets/:walletId/transactions/:id | ดูรายการเดียว |
| PUT | /wallets/:walletId/transactions/:id | แก้ไขรายการ |
| DELETE | /wallets/:walletId/transactions/:id | ลบรายการ |
| GET | /assets | ค้นหา/ดูสินทรัพย์ทั้งหมด |
| POST | /assets | เพิ่มสินทรัพย์ใหม่ (Find-or-Create) |
| GET | /wallets/:walletId/investment-transactions | ดูประวัติซื้อ-ขาย |
| POST | /wallets/:walletId/investment-transactions | บันทึกซื้อ/ขาย |
| DELETE | /wallets/:walletId/investment-transactions/:id | ลบรายการซื้อ-ขาย |
| GET | /wallets/:walletId/investment-holdings | ดูยอดถือครองปัจจุบัน |

---

## 👛 Wallets

### GET /wallets
**Auth required:** Yes
**Description:** ดูกระเป๋าเงินทั้งหมดของ user ที่ Login อยู่
**Response 200:**
```json
[
  { "id": 1, "name": "เงินเก็บฉุกเฉิน", "wallet_type": "cash", "created_at": "2026-08-01T10:00:00Z" }
]
```

### POST /wallets
**Auth required:** Yes
**Description:** สร้างกระเป๋าใหม่
**Request Body:**
```json
{ "name": "เงินเก็บฉุกเฉิน", "wallet_type": "cash" }
```
**Response 201:**
```json
{ "id": 1, "name": "เงินเก็บฉุกเฉิน", "wallet_type": "cash" }
```
**Errors:**
| Code | Reason |
|---|---|
| 400 | ชื่อซ้ำกับกระเป๋าอื่นของ user เดียวกัน / wallet_type ไม่ใช่ cash หรือ investment |

### GET /wallets/:id
**Auth required:** Yes
**Description:** ดูรายละเอียดกระเป๋าเดียว
**Response 200:** (object เดียวกับด้านบน)
**Errors:** `403` ไม่ใช่เจ้าของ, `404` ไม่พบ

### PUT /wallets/:id
**Auth required:** Yes
**Description:** แก้ไขชื่อและ/หรือประเภทกระเป๋า
**Request Body:**
```json
{ "name": "เงินเก็บระยะยาว", "wallet_type": "investment" }
```
**Response 200:** wallet ที่อัปเดตแล้ว
**Errors:**
| Code | Reason |
|---|---|
| 400 | ข้อมูลไม่ถูกต้อง |
| 403 | ไม่ใช่เจ้าของ |
| 404 | ไม่พบ |
| **409** | **พยายามเปลี่ยน wallet_type ทั้งที่กระเป๋านี้มี Transaction/Investment อยู่แล้ว — ต้องกระเป๋าว่างเปล่าเท่านั้นถึงเปลี่ยนประเภทได้** |

### DELETE /wallets/:id
**Auth required:** Yes
**Description:** ลบกระเป๋า — transactions/investment ข้างในถูกลบตาม (ON DELETE CASCADE)
**Response:** `204 No Content`
**Errors:** `403`, `404`

### GET /wallets/:id/balance
**Auth required:** Yes
**Description:** ยอดคงเหลือ คำนวณสดเสมอ พฤติกรรมต่างกันตาม `wallet_type`:
- `cash` → `SUM(transactions)` แยกบวก/ลบตาม transaction_type
- `investment` → ผลรวมต้นทุน (`quantity × avg_cost_per_unit`) ของทุก holding — **นี่คือ "ต้นทุนรวม" ไม่ใช่มูลค่าตลาดปัจจุบัน** เพราะไม่มีการดึงราคาสด

**Response 200:**
```json
{ "wallet_id": 1, "balance": 15000.00 }
```

---

## 💸 Transactions

### GET /wallets/:walletId/transactions
**Auth required:** Yes
**Description:** ดูรายการรายรับ-รายจ่ายของกระเป๋านี้
**Response 200:** array ของ transaction

### POST /wallets/:walletId/transactions
**Auth required:** Yes
**Description:** บันทึกรายรับ/รายจ่ายใหม่ — ใช้ได้เฉพาะกระเป๋าที่ `wallet_type = cash`
**Request Body:**
```json
{ "transaction_type": "expense", "amount": 250.00, "note": "ค่าข้าวเที่ยง", "transaction_date": "2026-08-20" }
```
**Response 201:** transaction ที่สร้าง
**Errors:**
| Code | Reason |
|---|---|
| 400 | amount ≤ 0 / กระเป๋านี้เป็น investment type |
| 403 / 404 | ไม่ใช่เจ้าของ / ไม่พบกระเป๋า |

### GET /wallets/:walletId/transactions/:id
**Auth required:** Yes — ดูรายการเดียว

### PUT /wallets/:walletId/transactions/:id
**Auth required:** Yes — แก้ไขรายการ (amount, note, transaction_date, transaction_type)

### DELETE /wallets/:walletId/transactions/:id
**Auth required:** Yes — ลบรายการ → `204 No Content`

---

## 🪙 Assets

### GET /assets
**Auth required:** Yes
**Description:** ดู/ค้นหาสินทรัพย์ทั้งหมด (ตารางกลาง ใช้ร่วมกันทุก user) — รองรับ `?search=BTC`
**Response 200:** array

### POST /assets
**Auth required:** Yes
**Description:** เพิ่มสินทรัพย์ใหม่ (Find-or-Create) — ถ้า `symbol` มีอยู่แล้ว คืนตัวเดิม ไม่สร้างซ้ำ
**Request Body:**
```json
{ "symbol": "BTC", "asset_type": "crypto" }
```
**Response 200/201:**
```json
{ "id": 3, "symbol": "BTC", "asset_type": "crypto" }
```
**Errors:** `400` asset_type ไม่ใช่ crypto หรือ stock

---

## 📈 Investment Transactions

### GET /wallets/:walletId/investment-transactions
**Auth required:** Yes — ดูประวัติซื้อ-ขายของกระเป๋านี้

### POST /wallets/:walletId/investment-transactions
**Auth required:** Yes
**Description:** บันทึกซื้อ/ขาย — ใช้ได้เฉพาะกระเป๋า `wallet_type = investment` และ `asset_id` ต้องมีอยู่แล้ว (สร้างผ่าน `POST /assets` ก่อน)

> ⚠️ **สำคัญมาก:** Handler นี้ต้อง Update `investment_holdings` (recompute `quantity` + `avg_cost_per_unit`) ในการ Database Transaction เดียวกันกับการ Insert เสมอ — ถ้าฝั่งใดฝั่งหนึ่งพลาดต้อง Rollback ทั้งคู่ ไม่งั้น Holdings จะไม่ตรงกับประวัติจริง

**Request Body:**
```json
{ "asset_id": 3, "transaction_type": "buy", "quantity": 0.05, "price_per_unit": 1500000, "transaction_date": "2026-08-20" }
```
**Response 201:** investment_transaction ที่สร้าง
**Errors:**
| Code | Reason |
|---|---|
| 400 | กระเป๋านี้เป็น cash type / sell มากกว่าจำนวนที่ถืออยู่จริง |
| 404 | ไม่พบ asset_id นี้ในระบบ |

### DELETE /wallets/:walletId/investment-transactions/:id
**Auth required:** Yes
**Description:** ลบรายการซื้อ-ขาย — อนุญาตให้ลบจริง

> ⚠️ ต้อง **Recompute `investment_holdings` ใหม่ทั้งหมด** สำหรับคู่ wallet+asset นั้น โดยไล่คำนวณ Average Cost ใหม่จากรายการที่เหลือทั้งหมดตามลำดับเวลา (ไม่ใช่แค่ลบแถวเฉยๆ) — เป็น Logic ที่ซับซ้อนกว่า CRUD ทั่วไป วางแผนเขียนเป็นฟังก์ชันแยกต่างหาก

**Response:** `204 No Content`

---

## 📊 Investment Holdings *(Read-only)*

### GET /wallets/:walletId/investment-holdings
**Auth required:** Yes
**Description:** ดูยอดถือครองปัจจุบันของกระเป๋านี้ — **ไม่มี POST/PUT/DELETE ตรงๆ** เพราะข้อมูลถูกดูแลผ่าน `investment_transactions` เท่านั้น
**Response 200:**
```json
[
  { "asset_id": 3, "symbol": "BTC", "quantity": 0.05, "avg_cost_per_unit": 1500000 }
]
```
