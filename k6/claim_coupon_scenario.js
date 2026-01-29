import http from 'k6/http';
import { check, sleep } from 'k6';
import { Counter } from 'k6/metrics';


/// base port on use
const BASE_URL = 'http://127.0.0.1:3000';
const COUPON_NAME_DOBLE_DIP = 'COUPON_DOUBLE_DIP';
const COUPON_NAME_FLASH_SALE = 'COUPON_FLASH_SALE';


const successClaims = new Counter('success_claims');
const outOfStockClaims = new Counter('out_of_stock_claims');
const duplicateClaims = new Counter('duplicate_claims');
const notFoundClaims = new Counter('not_found_claims');
const internalError = new Counter('internal_error_claims');


export const options = {
    scenarios: {
        flash_sale_attack: {
            executor: 'per-vu-iterations',
            vus: 50,
            iterations: 1,
            exec: 'flashSale',
        },

        double_dip_attack: {
            executor: 'per-vu-iterations',
            vus: 10,
            iterations: 1,
            exec: 'doubleDip',
            startTime: '5s',
        },
    },
};


export function flashSale() {
    const userId = `user-fs-${__VU}`;

    const payload = JSON.stringify({
        coupon_name: COUPON_NAME_FLASH_SALE,
        user_id: userId,
    });

    const params = {
        headers: { 'Content-Type': 'application/json' },
    };

    const res = http.post(`${BASE_URL}/api/coupons/claim`, payload, params);

    const ok = check(res, {
        'status is 201 or 400 or 404 or 409 or 500': (r) => r.status === 201 || r.status === 400 || r.status === 404 || r.status === 409 || r.status === 500
    });

    switch (res.status) {
        case 201:
            successClaims.add(1);
            break;
        case 400:
            outOfStockClaims.add(1);
            break;
        case 404:
            notFoundClaims.add(1);
            break;
        case 409:
            duplicateClaims.add(1);
            break;
        default:
            internalError.add(1);
    }

    sleep(0.1);
}



export function doubleDip() {

    const payload = JSON.stringify({
        coupon_name: COUPON_NAME_DOBLE_DIP,
        user_id: 'user-dd-1',
    });

    const params = {
        headers: { 'Content-Type': 'application/json' },
    };

    const res = http.post(`${BASE_URL}/api/coupons/claim`, payload, params);

    const ok = check(res, {
        'status is 201 or 400 or 404 or 409 or 500': (r) => r.status === 201 || r.status === 400 || r.status === 404 || r.status === 409 || r.status === 500
    });

    switch (res.status) {
        case 201:
            successClaims.add(1);
            break;
        case 400:
            outOfStockClaims.add(1);
            break;
        case 404:
            notFoundClaims.add(1);
            break;
        case 409:
            duplicateClaims.add(1);
            break;
        default:
            internalError.add(1);
    }

    sleep(0.1);
}
