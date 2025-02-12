/**
 * 라우팅 및 전체 구조
 */

import GlobalStyle from './components/GlobalStyle';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Header from './components/Header';
import Home from './pages/Home/Home';
import DIDPage from './pages/DID/DIDPage';
import VCPage from './pages/VC/VCPage';
import VPPage from './pages/VP/VPPage';
import Footer from './components/Footer';

const App = () => {
    return(
        <>
            <GlobalStyle/>

            <Router>
                <Header />
                
                <Routes>
                    <Route path='/' element={<Home/>} />
                    <Route path='/did' element={<DIDPage/>} />
                    <Route path='/vc' element={<VCPage/>} />
                    <Route path='/vp' element={<VPPage/>} />
                </Routes>

                <Footer/>
            </Router>
        </>
    );
};

export default App;