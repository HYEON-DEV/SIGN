// import React, { useState } from 'react';
import styles from '../styles/VPManager.module.scss';

const VPManager = () => {

    // const [vp, setVp] = useState('');
    // const handleCreateVp = async () => {
        // if (!selectedVc) {
        //     alert('VC를 선택하세요.');
        //     return;
        // }
        // VP 생성 API 호출
    // };

    return (
        <div className={styles.container}>
            <h2>VP 생성</h2>
            {/* <button className={styles.button} onClick={handleCreateVp} disabled={!selectedVc}>VP 생성</button>
            {vp && <p className={styles.vpDisplay}>생성된 VP: {vp}</p>} */}
        </div>
    );
};

export default VPManager;