<template>
  <div>

      <div class="panel">
        <panel-title :title="$route.meta.title"></panel-title>
        <div class="panel-body"
             v-loading="load_data"
             element-loading-text="拼命加载中">
          <el-row>
            <el-col :span="20">
              <el-form label-width="100px">

                <div class="panel-body">
                  <el-form-item style="font-weight:bold" label="发布环境:" label-width="100px">
                    <el-select v-model="env_id" filterable @change="select_env" placeholder="请选择" style="width: 200px;">
                      <el-option
                        v-for="item in envs"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value">
                        <span style="float: left">{{ item.label }}</span>
                        <span style="float: right; color: #8492a6; font-size: 13px">{{ item.lockstatus }}</span>
                      </el-option>
                    </el-select>
                    
                  </el-form-item>
                  <el-form-item style="font-weight:bold" label="项目名称:" label-width="100px">
                    <el-select v-model="pro_id" filterable @change="select_project" placeholder="请选择" style="width: 400px;">
                      <el-option
                        v-for="item in options"
                        :key="item.value"
                        :label="item.label"
                        :value="item.value">
                        <span style="float: left">{{ item.label }}</span>
                        <span style="float: right; color: #8492a6; font-size: 13px">{{ item.lockstatus }}</span>
                      </el-option>
                    </el-select>
                    <el-button @click.stop="on_refresh" size="small">
                      <i class="fa fa-refresh"></i>
                    </el-button>
                    
                  </el-form-item>

                </div>
                <el-form-item>
                  <el-button type="primary" @click="on_submit_form" :loading="on_submit_loading" :disabled="btn_submit_disable">创建
                  </el-button>
                  <el-button @click="add_lock" :loading="on_submit_loading" :disabled="btn_add_lock_disable">锁定</el-button>
                  <el-button @click="free_lock" :loading="on_submit_loading" :disabled="btn_free_lock_disable">解锁</el-button>
                  
                  <el-button @click="$router.back()">返回</el-button>
                </el-form-item>
              </el-form>
            </el-col>
          </el-row>
        </div>
      </div>
  </div>
</template>
<script type="text/javascript">
  import {panelTitle} from 'components'
  import {port_conf, port_code} from 'common/port_uri'
  import {tools_verify} from 'common/tools'
  import store from 'store'


  export default {
    data() {
      return {
        projects: null,
        options: [],
        pro_id: null,
        env_id: null,
        load_data: false,
        on_submit_loading: false,
        btn_submit_disable: true,
        btn_add_lock_disable: true,
        btn_free_lock_disable: true,
        user_role: store.state.user_info.user.Role,
        envs: []
      }
    },
    created() {
      // this.get_project_data()
      if (this.user_role == 1 || this.user_role == 20) {
        this.envs = [
          {value: 1, label: "测试环境"},
          {value: 2, label: "预发布环境"},
          {value: 3, label: "生产环境"}
        ]
      }else{
        this.envs = [{value: 2, label: "预发布环境"}]
      }

    },
    methods: {
      //获取数据
      get_project_data() {
        this.load_data = true
        this.$http.get(port_conf.mylist, {
          params: {
            env_id: this.env_id
          }
        })
          .then(({data: {data}}) => {
            var opData = []
            var uid = store.state.user_info.user.Id
            for (var i in data.table_data) {
              var value = data.table_data[i].id
              var env = ""
              var lockstatus = ""
              if(data.table_data[i].user_lock > 0){
                if(data.table_data[i].user_lock == uid){
                  lockstatus = data.table_data[i].lockuser+"锁定中"
                }else{
                  lockstatus = "锁定中"
                }
              }


              var env
                if(this.user_role == 1 || this.user_role == 20){
                   env = this.envs[this.env_id-1].label
                } else{
                  env = this.envs[0].label
                }
                var lable = env + "-" + data.table_data[i].name
                opData.push({label: lable, value: value, lockstatus:lockstatus})
              

            }
            this.projects = data.table_data
            this.options = opData
            this.load_data = false
            this.select_project()
            if (this.$route.query.id) {
              this.pro_id=this.$route.query.id
              this.on_submit_form()
            }
          }).catch(() => {

          this.load_data = false
        })
      },
      add_lock(){
        var proId = 0
        
          proId = this.pro_id
        
        if (proId) {
          this.$http.get(port_conf.lock, {
                  params: {
                    projectId:proId,
                    act:1
                  }
              })
              .then(({data: {data}}) => {
              this.$message({
                    message: "锁定成功！",
                    type: 'success'
                })
              this.get_project_data()
          })
        }
      },
      free_lock(){
        var proId = 0
        
        proId = this.pro_id
        
        if (proId) {
          this.$http.get(port_conf.lock, {
                  params: {
                    projectId:proId,
                    act:0
                  }
              })
              .then(({data: {data}}) => {
              this.$message({
                    message: "解锁成功！",
                    type: 'success'
                })
              this.get_project_data()
          })
        }
      },
      select_project(){
        var uid = store.state.user_info.user.Id
        var role = store.state.user_info.user.Role
        var proId = 0
        
        proId = this.pro_id
        if(!proId){
          return 
        }

        for (var i in this.projects){
          var project = this.projects[i]
          if(project.id == proId){
            if(project.user_lock > 0){
              if(project.user_lock == uid){
                this.btn_submit_disable=false
                this.btn_add_lock_disable=true
                this.btn_free_lock_disable=false
              }else{
                this.btn_submit_disable=true
                this.btn_add_lock_disable=true
                if(role == 1){
                  this.btn_free_lock_disable=false
                }else{
                  this.btn_free_lock_disable=true
                }
              }
            }else{
              this.btn_submit_disable=false
              this.btn_add_lock_disable=false
              this.btn_free_lock_disable=true
            }
          }
        }
      },
      select_env() {
        this.pro_id = null
        this.get_project_data()
      },
      on_refresh(){
        this.get_project_data()
      },
      //提交
      on_submit_form() {
        var proId = 0
        proId = this.pro_id
        
        if (proId) {
          for (var i in  this.projects) {
            var project = this.projects[i]
            if (proId == project.id) {
              console.log(project.repo_type )
              if (project.repo_type == "git") {
                this.$router.push({
                  name: 'taskGit',
                  query: {id: proId}
                })
                return
              }
              if (project.repo_type == "file") {
                this.$router.push({
                  name: 'taskFile',
                  query: {id: proId}
                })
                return
              }
                if (project.repo_type == "jenkins") {
                  this.$router.push({
                    name: 'taskJenkins',
                    query: {id: proId}
                  })
                  return
                }
            }
          }

        }
      }
    },
    components: {
      panelTitle
    }
  }
</script>
