// DID 생성 컴포넌트

// import {useState} from 'react';
// import {createDID} from '../api';
import styles from '../stylesDIDManager.module.scss';

const DIDManager = () => {
    // const [did, setDid] = useState('');

    // const handleCreateDID = async() => {    // 버튼 클릭 시 호출 
    //     try {
    //         const response = await createDID();
    //         setDid(response.data.did);
    //     } catch(error) {
    //         console.error('DID 생성 실패', error);
    //     }
    // };

    return (
        <div className={styles.container}>
            <h2> DID 생성 </h2>
            <button className={styles.button}> DID 생성 </button>
            {/* {did && <p className={styles.didDisplay}>생성된 DID: {did} </p>} */}
        </div>
    );
};

export default DIDManager;