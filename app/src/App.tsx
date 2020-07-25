import React from 'react';
import PostIcon from '@material-ui/icons/Book';
import UserIcon from '@material-ui/icons/Group';
import { Admin, Resource } from 'react-admin';
import jsonServerProvider from 'ra-data-json-server';
import Dashboard from './Dashboard';
import { PostList, PostEdit, PostCreate } from './Posts';
import { UserList } from './Users';

const dataProvider = jsonServerProvider("https://api-ten-minutes.lotteryjs.com");
const Title = () => (<div>Golang ❤️ MongoDB ❤️ React</div>)

const App = () => (
  <Admin title={<Title/>} dashboard={Dashboard} dataProvider={dataProvider}>
      <Resource name="posts" list={PostList} edit={PostEdit} create={PostCreate} icon={PostIcon} />
      <Resource name="users" list={UserList} icon={UserIcon} />
  </Admin>
)

export default App;
