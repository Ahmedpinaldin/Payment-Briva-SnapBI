package controller

import (
	"github.com/gin-gonic/gin"
) 

func ChangePassword(c *gin.Context) {   

	// IdUser := c.Params.ByName("id") 
	// form := &entity.ChangePassword{
	// 	PasswordNew: c.Request.FormValue("passwordNew"),
	// 	PasswordConfirm: c.Request.FormValue("passwordConfirm"),
	// 	PasswordOld: c.Request.FormValue("passwordOld"),
	// }

	// // Check Old Password
	// $post = $this->input->post(null, true);
	// $query = $this->db->query("SELECT username, password FROM user_login WHERE user_id ='$post[user_id]' AND password != '$post[password_old]'  ");
	// if ($query->num_rows() > 0) {
	// 	$this->form_validation->set_message('password_check', '*{field} salah');
	// 	return false;
	// } else {
	// 	return true;
	// }

	// // Change Password
	// $passowrd = $this->input->post('password_new');

	// $data = [
	// 	'password'  => $passowrd
	// ];
	// $where = "user_id = $_POST[user_id]";
	// $query = $this->db->update('user_login', $data, $where);
	// return $query;


	// IdQuery := c.Params.ByName("id") 
	// Id :=  "'" + IdQuery +"'"
	// fmt.Println(Id) 
	
	// db, err := config.NewConnection()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

	// query := "SELECT unit.nama AS nama_unit, unit.kode_unit AS kode_unit, unit.inisial AS inisial_unit, unit.alamat AS alamat_unit, cabang.inisial_cab AS inisial_cabang, cabang.nama AS nama_cabang, cabang.alamat AS alamat_cabang FROM unit LEFT JOIN cabang ON unit.inisial_cab = cabang.inisial_cab  WHERE unit.kode_unit = " + Id
	// resp, err := db.Query(query, Id)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, "Error creating  HTTP request")
    // }
    // defer resp.Close() 

	// var results []entity.ResponseProfile
	
	// for resp.Next() {
	// 	var result entity.ResponseProfile
	// 	err = resp.Scan(
	// 		&result.Nama_Unit,
	// 		&result.Kode_Unit,
	// 		&result.Inisial,
	// 		&result.Alamat_Unit,
	// 		&result.Inisial_Cabang,
	// 		&result.Nama_Cabang,
	// 		&result.Alamat_Cabang,
	// 	)
	// 	if err != nil {
	// 		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, nil)
	// 		return
	// 	}
	// 	results = append(results, result)
	// }
	// if err := resp.Err(); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, nil)
	// 	return
	// }
	// res := helper.BuildResponse(200, "Success", results)
	// c.AbortWithStatusJSON(http.StatusOK, res)
	
}																																																