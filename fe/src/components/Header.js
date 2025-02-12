/**
 * 헤더
 */

import styled from "styled-components";

const HeaderContainer = styled.header`
    padding: 10px;
    background-origin: var(--primary-color);
    /* color: white; */
    text-align: center;
`;

const Header = () => {
    return(
        <HeaderContainer>
            <h1> DID </h1>
        </HeaderContainer>
    );
};

export default Header;