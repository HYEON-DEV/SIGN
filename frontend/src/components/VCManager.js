// import { useEffect } from 'react';
import styles from '../styles/VCManager.module.scss';

// const VCManager = ( {onSelect} ) => {
    const VCManager = () => {

    // const [vcList, setVcList] = useState([]);
    // useEffect( () => {
    //     // VC 목록 조회 API 호출
    // }, [] );

    return (
        <div className={styles.container}>
            <h2> VC 목록 </h2>
            {/* <ul className={styles.vcList}>
                {vcList.map( (vc, index) => {
                    <li key={index} className={styles.vcItem} onClick={() => onSelect(vc) }>
                    </li>
                ))}
            </ul> */}
        </div>
    );
};

export default VCManager;